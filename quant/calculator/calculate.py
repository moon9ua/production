import pandas as pd
import numpy as np
from datetime import datetime, timedelta, timezone
from dateutil.relativedelta import relativedelta
import swifter
import time
import itertools


class DataLoad:
    def __init__(self):
        self.com_np, self.fin_np, self.pri_np = self.bring_datas()

    def bring_datas(self):
        print("getting datas...")
        com_df = pd.read_csv("./data/company.csv")
        fin_df = pd.read_csv("./data/finance.csv", encoding='cp949')
        pri_df = pd.read_csv("./data/price_monthly.csv")

        com_df = com_df[["ID", "MainSector"]]
        fin_df["Quarter"] = fin_df["Quarter"].swifter.apply(lambda x: datetime.strptime(x, "%Y-%m-%d"))
        pri_df["Date"] = pri_df["Date"].swifter.apply(lambda x: datetime.strptime(x, "%Y-%m-%d"))

        com_np = com_df.to_numpy()
        fin_np = fin_df.to_numpy()
        pri_np = pri_df.to_numpy()
        print("complete!")

        return com_np, fin_np, pri_np


class FindCode(DataLoad):
    def __init__(self):
        super().__init__()
        self.ref_date = []

    def trunc_dt_31(self, some_date):
        return datetime(some_date.year - 1, 12, 1, 0, 0, 0, 0)

    def rebalacing(self, start, end):
        if end is None:
            end = datetime(2021, 3, 31, 0, 0, 0, 0, tzinfo=timezone.utc)
        while True:
            temp = end
            end -= relativedelta(years=1)
            if end < start:
                break
            self.ref_date.append([end, temp])

    def apply_company(self, main_sector):
        com_np = self.com_np
        if len(main_sector) == 0:
            conditions = True
        else:
            conditions = False
            for ms in main_sector:
                conditions = conditions | (com_np[:, 1] == ms)

        com_np = com_np[conditions]
        return com_np[:, 0]

    def apply_each_finance(self, ref_date, fin_np, cond):
        fin_np = fin_np[fin_np[:, 3] == self.trunc_dt_31(ref_date[0]).date()]

        cond_idx = dict(
            net_rev=4,
            net_rev_r=5,
            net_prf=6,
            net_prf_r=7,
            de_r=8,
            per=9,
            psr=10,
            pbr=11,
            pcr=12,
            op_act=13,
            iv_act=14,
            fn_act=15,
            dv_yld=16,
            dv_pay_r=17,
            roa=18,
            roe=19
        )
        conditions = True

        for k, v in cond_idx.items():
            cond_arr = cond[str(k)]
            if cond_arr["min"] is None and cond_arr["max"] is None:
                continue
            else:
                conditions = conditions & (fin_np[:, v] >= cond_arr["min"]) & (fin_np[:, v] <= cond_arr["max"])

        fin_np = fin_np[conditions]

        return [fin_np[:, 0], ref_date]

    def apply_finance(self, cond, conditions):
        fin_np = self.fin_np

        mask = np.in1d(fin_np[:, 0], conditions)
        fin_np = fin_np[mask]
        conditions = [self.apply_each_finance(rf, fin_np, cond) for rf in self.ref_date]
        return conditions

    def apply_conditions(self, cond):
        self.rebalacing(cond["start_date"], cond["end_date"])
        conditions = self.apply_company(cond["main_sector"])
        conditions = self.apply_finance(cond, conditions)
        return conditions[::-1]


class Calculate(DataLoad):
    def __init__(self):
        super().__init__()
        self.prf = 0

    def trunc_dt_401(self, some_date):
        return datetime(some_date.year, some_date.month + 1, 1, 0, 0, 0, 0)

    def trunc_dt_430(self, some_date):
        return datetime(some_date.year, some_date.month + 1, 30, 0, 0, 0, 0)

    def get_max_loss_rate(self, ret):
        return ((1 + min(ret) / 100) / (1 + max(ret) / 100) - 1) * 100

    def get_holdings_count(self, code_list):
        return [len(c[0]) for c in code_list]

    def get_winning_percentage(self, ret):
        wr = [True if ret[i + 1] - ret[i] > 0 else False for i in range(len(ret) - 1)]
        return wr.count(True) / len(wr) * 100

    def get_annual_average_return(self, ret):
        av_c = [[r[i + 1] - r[i] for i in range(len(r) - 1)] for r in ret]
        av = [sum(a) / len(a) for a in av_c]
        return av

    def calculate_each_code(self, code, pri_np):
        profits = []
        pri_np = pri_np[(pri_np[:, 0] == code)]
        pri_np = pri_np[pri_np[:, 1].argsort()]  # sort by date
        self.prf = self.last_profit
        for i in range(len(pri_np) - 1):
            prf = (pri_np[i + 1, 5] / pri_np[i, 5]) - 1
            profit = (self.prf + 1) * (prf + 1) - 1
            self.prf = prf
            profits.append(profit * 100)
        if len(profits) != 12:
            return [0] * 12
        return profits

    def calculate_each_term(self, codes, ref_date):
        if len(codes) == 0:
            return [self.last_profit * 100.0] * 12
        profit = 0
        pri_np = self.pri_np
        pri_np = pri_np[(pri_np[:, 1] >= self.trunc_dt_401(ref_date[0]).date()) & \
                        (pri_np[:, 1] <= self.trunc_dt_430(ref_date[1]).date())]
        result = [self.calculate_each_code(c, pri_np) for c in codes]
        result_np = np.array(result)
        profit = result_np.sum(axis=0)
        profit /= len(codes)
        self.last_profit = profit[-1] / 100
        return profit

    def calculate_profit(self, code_list, cond):
        self.last_profit = 0
        acc_profit = [self.calculate_each_term(cl[0], cl[1]) for cl in code_list]
        annual_average_return = self.get_annual_average_return(acc_profit)
        chart = list(itertools.chain(*acc_profit))
        chart.insert(0, 0.0)
        winning_percentage = self.get_winning_percentage(chart)
        max_loss_rate = self.get_max_loss_rate(chart)
        holdings_count = self.get_holdings_count(code_list)
        return_dict = dict(
            cumulative_return=chart[-1],
            annual_average_return=annual_average_return[-1],
            winning_percentage=winning_percentage,
            max_loss_rate=max_loss_rate,
            holdings_count=holdings_count[-1],
            chart=dict(start_date=cond["start_date"], profit_rate_data=chart)
        )
        return return_dict


class QuantCalc(FindCode, Calculate):
    def __init__(self):
        super().__init__()

    def execute(self, **cond):
        code_list = self.apply_conditions(cond)
        return_dict = self.calculate_profit(code_list, cond)
        print(return_dict)
        return return_dict

<template>
  <div class="app-container">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="基本设置" name="first">基本设置</el-tab-pane>
      <el-tab-pane label="结算设置" name="second">
        <el-card shadow="hover" class="box-card">
          <el-form
            ref="dataForm"
            :model="form"
            label-position="left"
            label-width="150px"
            style="width: 800px; margin-left:50px;"
          >
            <el-form-item label="结算方式" prop="settlement_type">
              <el-radio-group v-model="form.settlement_type">
                <el-radio :label="1">分销佣金</el-radio>
                <el-radio :label="2">商城余额</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="提现方式" prop="withdrawalstype">
              <el-checkbox-group v-model="form.withdrawalstype">
                <el-checkbox label="1">商城余额</el-checkbox>
                <el-checkbox label="2">微信</el-checkbox>
                <el-checkbox label="3">支付宝</el-checkbox>
                <el-checkbox label="4">银行卡(自动提现)</el-checkbox>
                <el-checkbox label="5">银行卡(手动提现)</el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="佣金计算节点" prop="commission_calculation">
              <el-select v-model="form.commission_calculation" placeholder="请选择">
                <el-option
                  v-for="item in commission_calculations"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="佣金到账节点" prop="commission_arrival">
              <el-select v-model="form.commission_arrival" placeholder="请选择">
                <el-option
                  v-for="item in commission_arrivals"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="佣金提现免审核" prop="withdrawals_check">
              <el-switch v-model="form.withdrawals_check"></el-switch>
              <div class="help-block">开启后佣金提现佣金自动审核通过，无需手动操作</div>
            </el-form-item>
            <el-form-item label="佣金提现打款方式" prop="make_money">
              <el-radio-group v-model="form.make_money">
                <el-radio :label="1">自动</el-radio>
                <el-radio :label="2">手动</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="佣金个人所得税" prop="poundage">
              <el-input placeholder v-model="form.poundage">
                <template slot="append">%</template>
              </el-input>
            </el-form-item>
            <el-form-item label="佣金免打税区间" prop="withdrawalsbegin">
              <el-row>
                <el-col :span="10">
                  <el-input placeholder v-model="form.withdrawalsbegin" />
                </el-col>
                <el-col :span="14">
                  <el-input placeholder v-model="form.withdrawalsend">
                    <template slot="prepend">~</template>
                    <template slot="append">元</template>
                  </el-input>
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item label="佣金最低提现额度" prop="name">
              <el-input placeholder v-model="form.withdrawals_min">
                <template slot="append">元</template>
              </el-input>
            </el-form-item>
            <el-form-item label="佣金免审核提现额度" prop="name">
              <el-input placeholder v-model="form.withdrawals_cash">
                <template slot="append">元</template>
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="saveData">保存</el-button>
              <el-button>取消</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="推送通知" name="third">推送通知</el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { distributionSetting,getDistributionSetting } from "@/api/lime-admin/distributor";
export default {
  data() {
    return {
      activeName: "second",
      form: {
        settlement_type: 0,
        settlement_type: 0,
        withdrawalstype: [],
        commission_calculation:0,
        commission_arrival:0,
        withdrawals_check:0,
        make_money:0,
        poundage:0,
        withdrawalsbegin:0,
        withdrawalsend:0,
        withdrawals_min:0,
        withdrawals_cash:0,
        upgradeconditions: []
      },
      commission_calculations: [
        {
          value: "1",
          label: "实付款金额"
        },
        {
          value: "2",
          label: "小说漫画原价"
        },
        {
          value: "3",
          label: "小说漫画售价"
        },
        {
          value: "4",
          label: "小说漫画成本价"
        },
        {
          value: "5",
          label: "小说漫画利润价"
        }
      ],
      commission_arrivals: [
        {
          value: "1",
          label: "订单已经完成"
        }
      ],
    };
  },
  mounted() {
    this.getSetting();
  },
  methods: {
    handleClick(tab, event) {
      if (tab.name == "first") {
        this.$router.push({ path: "/distributor/distributionSetting" });
      }
    },
    saveData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          var data = { config_code: "settlementSetting", config_value: this.form };
          distributionSetting("settlementSetting", data).then(() => {
            this.$notify({
              title: "成功",
              message: "保存成功",
              type: "success",
              duration: 2000
            });
            this.$router.push({ path: "/distributor/settlementSetting" });
          });
        }
      });
    },
    async getSetting() {
      // 获取列表
      this.loading = true;
      try {
        const list = await getDistributionSetting('settlementSetting');
        if (list.data.result !== null){
          this.form = list.data.result;
        }
        
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
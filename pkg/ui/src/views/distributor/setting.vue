<template>
  <div class="app-container">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="基本设置" name="first">
        <el-card shadow="hover" class="box-card">
          <el-form
            ref="dataForm"
            :model="form"
            label-position="left"
            label-width="150px"
            style="width: 800px; margin-left:50px;"
          >
            <el-form-item label="小说漫画分销" prop="distribution_status">
              <el-switch v-model="form.distribution_status"></el-switch>
            </el-form-item>
            <el-form-item label="分销模式" prop="distribution_pattern">
              <el-select v-model="form.distribution_pattern" placeholder="请选择">
                <el-option
                  v-for="item in distribution_pattern_options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="分销内购" prop="purchasetype">
              <el-switch v-model="form.purchasetype"></el-switch>
              <div class="help-block">开启分销内购，分销商自己购买商品，可享受一级佣金，上级享受二级佣金，上上级享受三级佣金。</div>
            </el-form-item>
            <el-form-item label="店铺分销" prop="distribution_admin_status">
              <el-switch v-model="form.distribution_admin_status"></el-switch>
              <div class="help-block">开启后入驻店铺商品分销产生的佣金由店铺承担，开启前已发布的商品默认参与分销，请提醒入驻店铺编辑商品及时调整佣金比例。</div>
            </el-form-item>

            <el-form-item label="成为分销商条件" prop="distributorcondition">
              <el-radio-group v-model="form.distributorcondition">
                <el-radio :label="1">满足所有勾选条件</el-radio>
                <el-radio :label="2">满足勾选条件之一即可</el-radio>
                <el-radio :label="3">填写资料申请成为分销商</el-radio>
                <el-radio :label="4">无条件</el-radio>
              </el-radio-group>
              <el-card class="box-card">
                <el-checkbox-group v-model="form.distributorconditions">
                  <el-row>
                    <el-col :span="8">
                      <el-checkbox label="1" size="mini">消费金额达</el-checkbox>
                    </el-col>
                    <el-col :span="14">
                      <el-input placeholder v-model="form.pay_money">
                        <template slot="append">元</template>
                      </el-input>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="8">
                      <el-checkbox label="2" size="mini">订单数达</el-checkbox>
                    </el-col>
                    <el-col :span="14">
                      <el-input placeholder v-model="form.order_number">
                        <template slot="append">单</template>
                      </el-input>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="8">
                      <el-checkbox label="3" size="mini">购买商品，并完成订单</el-checkbox>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="8">
                      <el-checkbox label="4" size="mini">购买指定商品</el-checkbox>
                    </el-col>
                  </el-row>
                </el-checkbox-group>
              </el-card>
            </el-form-item>
            <el-form-item label="成为分销商下线条件" prop="lower_condition">
              <el-radio-group v-model="form.lower_condition">
                <el-radio :label="1">首次打开分享链接</el-radio>
                <el-radio :label="2">购买小说漫画，并付款</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="分销商自动审核" prop="distributorcheck">
              <el-switch v-model="form.distributorcheck"></el-switch>
              <div class="help-block">开启后无需手动审核分销商</div>
            </el-form-item>
            <el-form-item label="分销商必须完善资料" prop="distributorDatum">
              <el-switch v-model="form.distributorDatum"></el-switch>
              <div class="help-block">开启后未完善资料的分销商在提现时必须完善资料。</div>
            </el-form-item>
            <el-form-item label="是否开启跳级降级" prop="distributorgrade">
              <el-switch v-model="form.distributorgrade"></el-switch>
              <div class="help-block">开启后，分销商按照权重顺序跳级或跳降级，只要升级或降级条件满足则会一直执行。</div>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="saveData()">保存</el-button>
              <el-button>取消</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="结算设置" name="second">结算设置</el-tab-pane>
      <el-tab-pane label="申请协议" name="third">申请协议</el-tab-pane>
      <el-tab-pane label="文案样式" name="fourth">文案样式</el-tab-pane>
      <el-tab-pane label="推送通知" name="five">推送通知</el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { distributionSetting,getDistributionSetting } from "@/api/lime-admin/distributor";
export default {
  data() {
    return {
      activeName: "first",
      form: {
        distribution_status: 0,
        distribution_pattern: 1,
        purchasetype: 0,
        distribution_admin_status: 0,
        distributorcondition: 0,
        distributorconditions: [],
        lower_condition: 0,
        distributorcheck: 0,
        distributorDatum: 0,
        distributorgrade: 0
      },
      distribution_pattern_options: [
        {
          value: 1,
          label: "一级分销"
        },
        {
          value: 2,
          label: "二级分销"
        },
        {
          value: 3,
          label: "三级分销"
        }
      ],
    };
  },
  mounted() {
    this.getSetting();
  },
  methods: {
    handleClick(tab, event) {
      if (tab.name == "second") {
        this.$router.push({ path: "/distributor/settlementSetting" });
      }
    },
    saveData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          var data = { config_code: "distributionSetting", config_value: this.form };
          distributionSetting("distributionSetting", data).then(() => {
            this.$notify({
              title: "成功",
              message: "保存成功",
              type: "success",
              duration: 2000
            });
            this.$router.push({ path: "/distributor/distributionSetting" });
          });
        }
      });
    },
    async getSetting() {
      // 获取列表
      this.loading = true;
      try {
        const list = await getDistributionSetting('distributionSetting');
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
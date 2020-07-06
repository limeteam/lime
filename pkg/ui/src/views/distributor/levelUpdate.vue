<template>
  <el-card shadow="hover" class="box-card">
    <div>
      <el-row>
        <el-col :span="22">
          <span class="title"></span>
        </el-col>
        <el-col :span="2">
          <el-button type="primary" @click="onLevelList">等级列表</el-button>
        </el-col>
      </el-row>
      <el-divider>新增等级</el-divider>
    </div>
    <el-form
      ref="dataForm"
      :model="form"
      label-position="left"
      label-width="120px"
      style="width: 800px; margin-left:50px;"
    >
      <el-form-item label="等级名称" prop="name">
        <el-input v-model="form.name" style="width:50%" />
      </el-form-item>
      <el-form-item label="返佣类型" prop="recommendtype">
        <el-radio-group v-model="form.recommendtype">
          <el-radio :label="1">按比例返佣</el-radio>
          <el-radio :label="2">按固定返佣</el-radio>
        </el-radio-group>
        <div class="help-block">分销商下线购买商品后获得的分销返佣</div>
      </el-form-item>
      <el-form-item label="复购返佣" prop="buyagain_switch">
        <el-switch v-model="form.buyagain_switch" @change="changeBuyagain($event)" />
        <div class="help-block">开启后分销商下线再次购买商城商品时则按照复购返佣规则计算佣金</div>
        <div class="help-block" v-if="form.buyagain_switch == 0">分销商升级后，上级可获得推荐奖。0或空则不发放。</div>
      </el-form-item>
      <el-form-item label="返佣类型" prop="buyagain_recommendtype" v-if="form.buyagain_switch == 1">
        <el-radio-group v-model="form.buyagain_recommendtype">
          <el-radio :label="1">按比例返佣</el-radio>
          <el-radio :label="2">按固定返佣</el-radio>
        </el-radio-group>
        <div class="help-block">分销商下线购买商品后获得的分销返佣</div>
      </el-form-item>
      <el-form-item label="自动升级" prop="upgrade_switch">
        <el-switch v-model="form.upgrade_switch"></el-switch>
        <div class="help-block">开启后满足一定条件自动升级，关闭后则该等级不会自动升级。</div>
      </el-form-item>
      <el-form-item label="升级条件" prop="upgradecondition" v-if="form.upgrade_switch == 1">
        <el-radio-group v-model="form.upgradecondition">
          <el-radio :label="1">按比例返佣</el-radio>
          <el-radio :label="2">按固定返佣</el-radio>
        </el-radio-group>
        <el-card class="box-card">
          <el-checkbox-group v-model="form.upgradeconditions">
            <el-row>
              <el-col :span="8">
                <el-checkbox label="1" size="mini">一级分销商达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.number1">
                  <template slot="append">人</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="2" size="mini">二级分销商达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.number2">
                  <template slot="append">人</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="3" size="mini">三级分销商达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.number3">
                  <template slot="append">人</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="4" size="mini">团队人数达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.number4">
                  <template slot="append">人(分销商)</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="5" size="mini">下线客户数达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.number5">
                  <template slot="append">人(非分销商)</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="6" size="mini">下线总人数达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.offline_number">
                  <template slot="append">人(分销商+非分销商)</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="7" size="mini">指定推荐等级</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-row>
                  <el-col :span="10">
                    <el-select v-model="form.upgrade_level" placeholder>
                      <el-option label="默认指定等级" value="1"></el-option>
                      <el-option label="等级1" value="2"></el-option>
                    </el-select>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="form.level_number">
                      <template slot="prepend">达</template>
                      <template slot="append">人</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="8" size="mini">分销订单金额达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.order_money">
                  <template slot="append">元</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="9" size="mini">分销订单达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.order_number">
                  <template slot="append">单</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="10" size="mini">内购订单金额达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.selforder_money">
                  <template slot="append">元</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="11" size="mini">内购订单达</el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="form.selforder_number">
                  <template slot="append">单</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="12" size="mini">购买指定商品</el-checkbox>
              </el-col>
            </el-row>
          </el-checkbox-group>
        </el-card>
      </el-form-item>

      <el-form-item label="自动降级" prop="downgrade_switch">
        <el-switch v-model="form.downgrade_switch"></el-switch>
        <div class="help-block">开启后不满足指定条件自动降级，关闭后则该等级不会自动降级。</div>
      </el-form-item>

      <el-form-item label="降级条件" prop="downgradecondition" v-if="form.downgrade_switch == 1">
        <el-radio-group v-model="form.downgradecondition">
          <el-radio :label="1">满足所有勾选条件</el-radio>
          <el-radio :label="2">满足勾选条件之一即可</el-radio>
        </el-radio-group>
        <el-card class="box-card">
          <el-checkbox-group v-model="form.downgradeconditions">
            <el-row>
              <el-col :span="8">
                <el-checkbox label="1" size="mini">团队订单量</el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-row>
                  <el-col :span="10">
                    <el-input placeholder v-model="form.team_number_day">
                      <template slot="append">天内</template>
                    </el-input>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="form.team_number">
                      <template slot="prepend">少于</template>
                      <template slot="append">单</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="2" size="mini">团队订单金额</el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-row>
                  <el-col :span="10">
                    <el-input placeholder v-model="form.team_money_day">
                      <template slot="append">天内</template>
                    </el-input>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="form.team_money">
                      <template slot="prepend">少于</template>
                      <template slot="append">单</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="3" size="mini">内购订单金额</el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-row>
                  <el-col :span="10">
                    <el-input placeholder v-model="form.self_money_day">
                      <template slot="append">天内</template>
                    </el-input>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="form.self_money">
                      <template slot="prepend">少于</template>
                      <template slot="append">单</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
          </el-checkbox-group>
        </el-card>
      </el-form-item>
      <el-form-item label="权重" prop="weight">
        <el-input-number v-model="form.weight" :step="1" step-strictly></el-input-number>
        <div class="help-block">等级权重，数字不能重复，数字越大等级越高。按设置的权重大小从低到高进行升级。</div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="updateData()">修改</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import {
  updateDistributorLevel,
  getDistributorLevel
} from "@/api/lime-admin/distributorLevel";
export default {
  name: "CreateDistributorLevel",
  data() {
    return {
      form: {
        name: "",
        recommendtype: 1,
        buyagain_switch: 1,
        auto_upgrade: 0,
        upgrade_switch: 0,
        upgradecondition: 0,
        upgradeconditions: [],
        number1: 0,
        number2: 0,
        number3: 0,
        number4: 0,
        number5: 0,
        offline_number: 0,
        upgrade_level: 0,
        level_number: 0,
        order_money: 0,
        order_number: 0,
        selforder_money: 0,
        selforder_number: 0,
        downgradecondition: 0,
        downgrade_switch: 0,
        upgrade_conditions: [],
        adaptive_degradation: 0,
        degradation_conditions: [],
        downgradeconditions: [],
        weight: 0
      }
    };
  },
  mounted() {
    this.getLevelInfo();
  },
  methods: {
    onLevelList() {
      this.$router.push({ path: "/distributor/distributorLevelList" });
    },
    resetForm() {
      // 重置
      this.$refs["dataForm"].resetFields();
    },
    changeBuyagain: function($event) {
      console.log(this.form.buyagain_switch);
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          var data = {
            name: this.form.name,
            recommendtype: this.form.recommendtype,
            buyagain_switch: this.form.buyagain_switch,
            auto_upgrade: this.form.auto_upgrade,
            upgrade_conditions: {
              upgradecondition: this.form.upgradecondition,
              upgradeconditions: this.form.upgradeconditions,
              number1: this.form.number1,
              number2: this.form.number2,
              number3: this.form.number3,
              number4: this.form.number4,
              number5: this.form.number5,
              offline_number: this.form.offline_number,
              level_number: this.form.level_number,
              order_money: this.form.order_money,
              order_number: this.form.order_number,
              selforder_money: this.form.selforder_money,
              selforder_number: this.form.selforder_number
            },
            adaptive_degradation: this.form.adaptive_degradation,
            degradation_conditions: {
              downgradeconditions: this.form.downgradeconditions,
              downgradecondition: this.form.downgradecondition,
              team_number_day: this.form.team_number_day,
              team_number: this.form.team_number,
              team_money_day: this.form.team_money_day,
              team_money: this.form.team_money,
              self_money_day: this.form.self_money_day,
              self_money: this.form.self_money
            },
            weight: this.form.weight
          };
          updateDistributorLevel(id, data).then(() => {
            this.dialogFormVisible = false;
            this.$notify({
              title: "成功",
              message: "新增成功",
              type: "success",
              duration: 2000
            });
            this.$store.dispatch("tagsView/delView", this.$route);
            this.$router.push({ path: "/distributor/distributorLevelList" });
          });
        }
      });
    },
    async getLevelInfo() {
      // 获取列表
      this.loading = true;
      try {
        var id = this.$route.query.id;
        const list = await getDistributorLevel(id);
        var data = list.data.result;
        this.form = {
          name: data.name,
          recommendtype: data.recommendtype,
          buyagain_switch: data.buyagain_switch,
          auto_upgrade: data.auto_upgrade,
          upgrade_switch: 0,
          upgradecondition: data.upgrade_conditions.upgrade_condition,
          upgradeconditions: data.upgrade_conditions.upgrade_conditions,
          number1: data.upgrade_conditions.number1,
          number2: data.upgrade_conditions.number2,
          number3: data.upgrade_conditions.number3,
          number4: data.upgrade_conditions.number4,
          number5: data.upgrade_conditions.number5,
          offline_number: data.upgrade_conditions.offline_number,
          upgrade_level: data.upgrade_conditions.upgrade_level,
          level_number: data.upgrade_conditions.level_number,
          order_money: data.upgrade_conditions.order_money,
          order_number: data.upgrade_conditions.order_number,
          selforder_money: data.upgrade_conditions.selforder_money,
          selforder_number: data.upgrade_conditions.selforder_number,
          downgradecondition: data.degradation_conditions.downgradecondition,
          downgrade_switch: data.degradation_conditions.downgrade_switch,
          upgrade_conditions: data.degradation_conditions.upgrade_conditions,
          adaptive_degradation: data.degradation_conditions.adaptive_degradation,
          degradation_conditions: data.degradation_conditions.degradation_conditions,
          downgradeconditions: data.degradation_conditions.downgradeconditions,
          weight: data.weight
        };
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
<style lang="scss" scoped>
.title {
  font-weight: 700 !important;
  font-size: 16px;
  padding: 20px 20px;
}
.help-block {
  display: block;
  margin-top: 5px;
  margin-bottom: 10px;
  color: #737373;
}
</style>
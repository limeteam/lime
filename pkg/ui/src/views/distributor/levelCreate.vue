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
        <el-input v-model="form.name" style="width:50%"/>
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
      <el-form-item label="自动升级" prop="auto_upgrade">
        <el-switch v-model="form.auto_upgrade"></el-switch>
        <div class="help-block">开启后满足一定条件自动升级，关闭后则该等级不会自动升级。</div>
      </el-form-item>
      <el-form-item label="升级条件" prop="auto_upgrade" v-if="form.auto_upgrade == 1">
        <el-radio-group v-model="form.recommendtype">
          <el-radio :label="1">按比例返佣</el-radio>
          <el-radio :label="2">按固定返佣</el-radio>
        </el-radio-group>
        <el-card class="box-card">
          <el-checkbox-group v-model="form.upgradeconditions">
            <el-row>
              <el-col :span="8">
                <el-checkbox label="一级分销商达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder="" v-model="input2">
                  <template slot="append">人</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="二级分销商达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder="" v-model="input2">
                  <template slot="append">人</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="三级分销商达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">人</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="团队人数达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">人(分销商)</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="下线客户数达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">人(非分销商)</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="下线总人数达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">人(分销商+非分销商)</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="指定推荐等级" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-row>
                  <el-col :span="10">
                    <el-select v-model="select" placeholder>
                  <el-option label="默认指定等级" value="1"></el-option>
                  <el-option label="等级1" value="2"></el-option>
                </el-select>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="input2">
                      <template slot="prepend">达</template>
                      <template slot="append">人</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="分销订单金额达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">元</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="分销订单达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">单</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="内购订单金额达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">元</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="内购订单达" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2">
                  <template slot="append">单</template>
                </el-input>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="购买指定商品" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="14">
                <el-input placeholder v-model="input2"></el-input>
              </el-col>
            </el-row>
          </el-checkbox-group>
        </el-card>
      </el-form-item>

      <el-form-item label="自动降级" prop="adaptive_degradation">
        <el-switch v-model="form.adaptive_degradation"></el-switch>
        <div class="help-block">开启后不满足指定条件自动降级，关闭后则该等级不会自动降级。</div>
      </el-form-item>

      <el-form-item label="降级条件" prop="auto_upgrade" v-if="form.adaptive_degradation == 1">
        <el-radio-group v-model="form.recommendtype">
          <el-radio :label="1">满足所有勾选条件</el-radio>
          <el-radio :label="2">满足勾选条件之一即可</el-radio>
        </el-radio-group>
        <el-card class="box-card">
          <el-checkbox-group v-model="form.upgradeconditions">
            <el-row>
              <el-col :span="8">
                <el-checkbox label="团队订单量" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-row>
                  <el-col :span="10">
                    <el-input placeholder v-model="input2">
                      <template slot="append">天内</template>
                    </el-input>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="input2">
                      <template slot="prepend">少于</template>
                      <template slot="append">单</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="团队订单金额" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-row>
                  <el-col :span="10">
                    <el-input placeholder v-model="input2">
                      <template slot="append">天内</template>
                    </el-input>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="input2">
                      <template slot="prepend">少于</template>
                      <template slot="append">单</template>
                    </el-input>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="8">
                <el-checkbox label="内购订单金额" name="type" size="mini"></el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-row>
                  <el-col :span="10">
                    <el-input placeholder v-model="input2">
                      <template slot="append">天内</template>
                    </el-input>
                  </el-col>
                  <el-col :span="14">
                    <el-input placeholder v-model="input2">
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
        <el-button type="primary" @click="createData">添加</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import { createDistributorLevel } from "@/api/lime-admin/distributorLevel";
export default {
  name: "CreateDistributorLevel",
  data() {
    return {
      form: {
        name: "",
        recommendtype: 1,
        buyagain_switch: 1,
        auto_upgrade: 0,
        upgradeconditions: [],
        upgrade_conditions: "",
        adaptive_degradation: 0,
        degradation_conditions: "",
        weight: ""
      }
    };
  },
  mounted() {},
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
    createData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          createDistributorLevel(this.form).then(() => {
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
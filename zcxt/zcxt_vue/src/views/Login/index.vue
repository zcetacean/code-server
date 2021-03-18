<template>
<div id="login">
  <div class="login-wrap">
    <div class="in">
    <div class="title block">综测系统后台管理</div>
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="login-form" size="medium">
      <el-form-item label="账号" prop="username"  class="item-form">
        
        <el-input type="text" v-model="ruleForm.username" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item label="密码" prop="password" class="item-form">
        
          <el-input type="password" v-model="ruleForm.password" autocomplete="off" maxlength="20" minlength="6"></el-input>
      </el-form-item>

      <el-form-item class="item-login">
          <el-button type="primary" @click="submitForm('ruleForm')" class="login-btn block">登录</el-button>
      </el-form-item>
    </el-form>
    </div>
  </div>
</div>
</template>

<script>
import { stripscript, validateEmail, validatePass, validateVCode } from '@/utils/validate.js';
export default {
    data() {
      //验证用户名
      var validateUsername = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入用户名'));
        } else if(validateEmail(value)){
            callback(new Error('用户名格式有误'))
        }else{
          callback(); 
        }
      };
      //验证密码
      var validatePassword = (rule, value, callback) => {
        // 过滤后的数据
        this.ruleForm.password = stripscript(value);
        // console.log(this.ruleForm.password,value);
        // value为用户输入的值
        if (value === '') {
          callback(new Error('请输入密码'));
        } else if (validatePass(value)) {
          callback(new Error('密码为6至20位数字+字母'));
        } else if (this.ruleForm.password != value) {
          callback(new Error('密码含有非法字符'));
        } else {
          callback();
        }
      };
      var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入账号'));
        } else {
          if (this.ruleForm.checkPass !== '') {
            this.$refs.ruleForm.validateField('checkPass');
          }
          callback();
        }
      };
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          callback();
        }
      };

      return {
        ruleForm: {
          pass: '',
          checkPass: '',
          age: ''
        },
        rules: {
          pass: [
            { validator: validatePass, trigger: 'blur' }
          ],
          checkPass: [
            { validator: validatePass2, trigger: 'blur' }
          ],
        }
      };
    },

    /********************************************************
       * 函数定义
       */
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$message({
            message: '登录成功',
            type: 'success'
          });
            this.$router.push({
            name: 'Student'
          });
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }


      /********************************************************
       * 声明数据
       */



  

      
      
      //提交表单
     
</script>

<style scoped lang="scss">
#login {
    height:100vh;
}
.login-wrap{
    width: 494px;
    margin: auto;
    padding-top: 13%;
    border-radius: 4px;
}
.in{
  padding: 25px 55px 20px 0px;
  margin-left:20px;
  border:"true";
  border-radius: 44px;
  box-shadow: 6px 10px 16px -5px #646464
}
.title{
  margin-left: 155px;
  font-size: 24px;
}
.login-form {
    margin-top: 40px;
    label {
        display: block;
        margin-bottom: 3px;
        font-size: 16px;
        color: #000000;
    }
    .item-form {
        margin-bottom: 13px;
    }
    .block {
        display: block;
        width: 100%;
    }
    .element.style {
    margin-left: 0px !important;
}
    .login-btn { 
      margin-top: 25px;
      }
}
.el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 4px;
  }
</style>
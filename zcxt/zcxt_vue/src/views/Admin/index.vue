<template>
<div class="student">
  <el-button-group style="float:right;margin-bottom:30px;">
  <el-button >导入</el-button>
  <el-button >导出</el-button>
  <el-button >删除</el-button>
  <el-button >更改</el-button>
</el-button-group>
    <el-table
    ref="multipleTable"
    :data="tableData"
    tooltip-effect="dark"
    style="width: 100%"
    >
    <el-table-column
      type="selection"
      width="55">
    </el-table-column>
    <el-table-column
      prop="name"
      label="名字"
      width="150">
    </el-table-column>
    <el-table-column
      prop="user"
      label="账号"
      width="200">
    </el-table-column>
    <el-table-column
      prop="password"
      label="密码"
      width="200">
    </el-table-column>
    <el-table-column
      prop="level"
      label="管理员等级"
      >
    </el-table-column>
  </el-table>
  <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage4"
      :page-sizes="[100, 200, 300, 400]"
      :page-size="100"
      layout="total, sizes, prev, pager, next, jumper"
      :total="400"
      style="float:left;margin-top:30px;">
    </el-pagination>
</div>
</template>

<script>
import{ reactive ,ref } from '@vue/composition-api';
export default {
    name: 'Student',
    setup(props){
    const tableData = reactive([{
          name: '杨泽钦',
          user: '6736589123',
          password: '123456',
          level: '二级管理员'
        },{
          name: '杨燕佳',
          user: '6736589123',
          password: '123456',
          level: '一级管理员'
        },{
          name: '杨泽钦',
          user: '6736589123',
          password: '123456',
          level: '一级管理员'
        },
        ])

    return {
        tableData
    }
    }
}
</script>

<style lang="scss" scoped>
.student{
  display:inline-block;
  margin-top: 20px;
  width: 950px;
  height: 900px;
}


</style>
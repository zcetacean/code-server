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
      prop="number"
      label="学号"
      width="170">
    </el-table-column>
    <el-table-column
      prop="state"
      label="状态"
      width="150">
    </el-table-column>
    <el-table-column
      prop="img"
      label="图片"
      width="200">
    </el-table-column>
    <el-table-column
      prop="class"
      label="请求分数"
      width="150">
    </el-table-column>
    <el-table-column
      prop="count"
      label="最终分数"
      show-overflow-tooltip
      width="150"
      >
    </el-table-column>
    <el-table-column
      prop="count"
      label="最高审核人"
      show-overflow-tooltip
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
          number: '2018031456118',
          state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
        }, {
          number: '2018031456118',
           state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
        }, {
          number: '2018031456118',
           state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
        }, {
          number: '2018031456118',
          state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
        }, {
          number: '2018031456118',
          state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
        }, {
          number: '2018031456118',
          state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
        }, {
          number: '2018031456118',
          state: '第一次审批',
          img: '123456',
          class: '18信管1班',
          count: '15'
    }])

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
  width: 1000px;
  height: 900px;
}


</style>
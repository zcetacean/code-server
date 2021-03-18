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
      prop="item"
      label="综测项目"
      width="200">
    </el-table-column>
    <el-table-column
      prop="score"
      label="应加分数"
      width="100">
    </el-table-column>
    <el-table-column
      prop="remark"
      label="备注说明"
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
          item: '院级三好学生',
          score: '10',
          remark: '校级和院级不同'
        },{
          item: '校级三好学生',
          score: '15',
          remark: '校级和院级不同'
        },{
          item: '优秀干部',
          score: '6',
          remark: ''
        }
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
<template>
<div class="center">
  <el-button-group style="float:right;margin:20px 0 20px;">
  <el-button >导入</el-button>
  <el-button >导出</el-button>
  <el-button >删除</el-button>
  <el-button >更改</el-button>
</el-button-group>
  <div class="custom-tree-container">
    <div class="block">
      <el-tree
        :data="data"
        show-checkbox
        node-key="id"
        default-expand-all
        :expand-on-click-node="false">
        <span class="custom-tree-node" slot-scope="{ node, data }">
          <span>{{ node.label }}</span>
          <span>
            <el-button
              type="text"
              size="mini"
              @click="() => append(data)">
              Append
            </el-button>
            <el-button
              type="text"
              size="mini"
              @click="() => remove(node, data)">
              Delete
            </el-button>
          </span>
        </span>
      </el-tree>
    </div>
</div>
</div>
</template>
<script>
import { ref, reactive, isRef, toRefs, onMounted, computed }  from "@vue/composition-api";
export default {
    name: 'Tree',
    setup(props,{ root }){
      const data = reactive([{
         id: 1,
        label: '一级 1',
        children: [{
          id: 4,
          label: '二级 1-1',
          children: [{
            id: 9,
            label: '三级 1-1-1'
          }, {
            id: 10,
            label: '三级 1-1-2'
          }]
        }]
      }, {
        id: 2,
        label: '一级 2',
        children: [{
          id: 5,
          label: '二级 2-1'
        }, {
          id: 6,
          label: '二级 2-2'
        }]
      }, {
        id: 3,
        label: '一级 3',
        children: [{
          id: 7,
          label: '二级 3-1'
        }, {
          id: 8,
          label: '二级 3-2'
        }]
      }])


      const append = (data) =>{
        const newChild = { id: id++, label: 'testtest', children: [] };
        if (!data.children) {
          root.$set(data, 'children', []);
        }
        data.children.push(newChild);
      }

      const remove = (node, data) =>{
        const parent = node.parent;
        const children = parent.data.children || parent.data;
        const index = children.findIndex(d => d.id === data.id);
        children.splice(index, 1);
      }


      return {
        data,
        append,
        remove
      }
    }
}
</script>
<style lang="scss" scoped>
.center{
  display: inline-block;
  width:500px;
}
.custom-tree-container{
  margin-top:90px;
}
.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
  }
</style>
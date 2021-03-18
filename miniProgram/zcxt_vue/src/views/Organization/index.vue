<template>
    <div class="center">
        <el-tabs :tab-position="tabPosition" style="height: 200px;">
            <el-tab-pane label="组织树">用户管理</el-tab-pane>
            <el-tab-pane label="综测条目">配置管理</el-tab-pane>
        </el-tabs>  
    </div>
</template>

<script>
import{ reactive ,ref } from '@vue/composition-api';
export default {
    name: 'Organization',
    setup(props){
        const tabPosition = ref("top");

        return {
        tabPosition
        }
    }
}
</script>

<style lang="scss" scoped>
.center{
    display:inline-block;
}
</style>
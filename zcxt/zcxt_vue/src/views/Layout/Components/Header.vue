<template>
    <div id="header">
        <div class="center">
        <el-menu 
        class="el-menu-demo" 
        mode="horizontal" 
        router>
            <el-menu-item index="/student">学生管理</el-menu-item>
            <el-menu-item index="/verify">综测管理</el-menu-item>
            <el-submenu index="/organization">
                <template slot="title">组织管理</template>
                <el-menu-item index="/tree">组织树</el-menu-item>
                <el-menu-item index="/item">综测条目</el-menu-item>
            </el-submenu>
            <el-menu-item index="/admin" >管理员</el-menu-item>
        </el-menu>
        </div>
    </div>
</template>

<script>
import{ reactive ,ref } from '@vue/composition-api';
export default {
        name:'index',
        setup(props, { root }){

            const routers = reactive(root.$router.options.routes);


            return {
            routers
        };
        }
}
</script>

<style lang="scss" scoped>
#header{
    text-align:center;  
    padding-top: 30px;
    height:65px;
    > .center{
        display:inline-block;
    }
}

</style>
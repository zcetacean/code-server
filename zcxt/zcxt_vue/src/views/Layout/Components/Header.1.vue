<template>
    <div id="header">
        <div class="center">
        <el-menu 
        class="el-menu-demo" 
        mode="horizontal" 
        router>
            <template v-for="(item) in routers">
                <el-menu-item v-if="!item.hidden" :key="item.id" :index="item.path">
                    {{ item.meta.name }}
                </el-menu-item>
            </template>
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
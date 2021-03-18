<template>
    <div>
        <LayoutHeader />
        <LayoutMain />
    </div>
</template>

<script>
import LayoutHeader from "./Components/Header";
import LayoutMain from "./Components/Main";
import{ reactive ,ref } from '@vue/composition-api';
export default {
        name:'index',
        components: { LayoutHeader, LayoutMain},
        setup(props){
            


            return {
            
        };
        }
}
</script>

<style lang="scss" scoped>



</style>
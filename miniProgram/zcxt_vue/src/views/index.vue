<template>
    <div>
        <LayoutHeader />
        <LayoutMain />
    </div>
</template>

<script>    
import LayoutHeader from './Tabs/index.vue';
import{ reactive ,ref } from '@vue/composition-api';
export default {
        name:'index',
        components: {Student, Item, Organization, Admin},
        setup(props){
            const activeName = ref('first');


            return {
            activeName
        };
        }
}
</script>

<style lang="scss" scoped>



</style>
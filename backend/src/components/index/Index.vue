<template>
    <div>
        首页
    </div>
</template>

<script>
import { BreadcrumbUtil } from "@/lib/BreadcrumbUtil";
import tableData from "@/data/tableData";

export default {
    data() {
        return {
            tableData
        }
    },
    components: {
    },
    created() {
    },
    methods: {
        getBreadcrumb: function() {
            let breadcrumbUtil = new BreadcrumbUtil();
            breadcrumbUtil.setItem("首页", "");

            return breadcrumbUtil.getBreadcrumbs();
        }
    }
}
</script>

<style scoped>
    .el-table td, .el-table th{
        padding: 0px!important;
    }
    .el-pagination {
        margin: 10px 0 0;
        text-align: right;
    }
</style>

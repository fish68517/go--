/**
 * 泡药管理
 * @author fengxh
 * @since 2025/1/13
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        ,form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'word_person', width: 250, title: '员工姓名', align: 'center'}
            , {field: 'word_content', width: 100, title: '工作内容', align: 'center'}
            , {field: 'dosage', width: 100, title: '贴数', align: 'center'}


        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("员工工作量统计");

    }
});

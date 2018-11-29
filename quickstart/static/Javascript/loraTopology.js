/**
 * Created by Administrator on 2018/10/1.
 */

//1.准备树的数据
/*var data = [
    {name:"广东",value:"01",sj:"-"},
    {name:"广州",value:"0101",sj:"01"},
    {name:"潮州",value:"0102",sj:"01"},
    {name:"深圳",value:"0103",sj:"01"},
    {name:"茂名",value:"0104",sj:"01"},
    {name:"揭阳",value:"0105",sj:"01"},
    {name:"萝岗",value:"010101",sj:"0101"},
    {name:"天河",value:"010102",sj:"0101"},
    {name:"黄埔",value:"010103",sj:"0101"},
    {name:"白云",value:"010104",sj:"0101"},
    {name:"枫溪",value:"010201",sj:"0102"},
    {name:"枫桥",value:"010202",sj:"0102"},
    {name:"罗湖",value:"010301",sj:"0103"}
];*/
var treeData ;
function setTopology() {
    $.get("/lora/topology",function(result){
        var data = JSON.parse(result);
        //result数据添加到容器中;
          treeData = transData(data, 'value', 'sj', 'children');
        //展示数据
         drawTree(treeData);
    });
}



/**2.数据处理成层级关系的数据***/
function transData(a, idStr, pidStr, childrenStr) {
    var r = [], hash = {}, id = idStr, pid = pidStr, children = childrenStr, i = 0, j = 0, len = a.length;
    for (; i < len; i++) {
        hash[a[i][id]] = a[i];
    }
    for (; j < len; j++) {
        var aVal = a[j], hashVP = hash[aVal[pid]];
        if (hashVP) {
            !hashVP[children] && (hashVP[children] = []);
            hashVP[children].push(aVal);
        } else {
            r.push(aVal);
        }
    }
    return r;
}

/**
 *3. 画树
 */
function drawTree(treeData) {
    var  myChart = echarts.init(document.getElementById("container"));//div元素节点的对象
    myChart.setOption({
        tooltip : {
            trigger : 'item',
            triggerOn : 'mousemove'
        },
        series : [ {
            type : 'tree',
            name : 'TREE_ECHARTS',
            data : treeData,
            top : '10%',
            left : '2%',
            bottom : '25%',
            right : '1%',
            symbolSize: [70, 30],
            orient: 'vertical',
            rootLocation: {x: 'center',y: 230},
            nodePadding: 20,
            layerPadding:40,
            symbol: 'rectangle',

            borderColor:'black',
            itemStyle: {
                normal: {
                    borderColor: '#006dcc',
                    color: 'white',//节点背景色
                    label: {
                        show: true,
                        borderWidth:0,
                        position: 'inside',
                        textStyle: {
                            color: 'black',
                            fontSize: 15
                        }
                    },
                    lineStyle: {
                        color: '#006dcc',
                        width: 1,
                        curveness:0,
                        type: 'solid'
                    }
                },
                emphasis: {
                    label: {
                        show: false,
                        color:'#006dcc'
                    }
                }
            },

            expandAndCollapse : true ,

            initialTreeDepth : 2  //展示层级数,默认是2
        } ]
    });
    //4.树绑定事件
    myChart.on('click', function(params) {
        var name = params.data.name;//点击的节点的name
        var value = params.data.value;//点击的节点的value
        //调用点击事件
        clickNode(name,value);
    });
}
//节点的点击事件
function clickNode(name,value){
   // alert(name+'--的值：'+value);
}

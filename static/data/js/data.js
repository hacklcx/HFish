window.onresize = function () {
    main1.resize();
    main2.resize();
    main3.resize();
    main4.resize();
    main5.resize();
    main6.resize();
    worldMap.resize();
};

var main1 = echarts.init(document.getElementById('main1'), 'dark');
var main2 = echarts.init(document.getElementById('main2'), 'dark');
var main3 = echarts.init(document.getElementById('main3'), 'dark');
var main4 = echarts.init(document.getElementById('main4'), 'dark');
var main5 = echarts.init(document.getElementById('main5'));
var main6 = echarts.init(document.getElementById('main6'));
var worldMap = echarts.init(document.getElementById('worldMap'));


function main1_data() {
    $.ajax({
        type: "GET",
        url: "/data/get/china",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var listx = [];

            for (var i = 0; i < d.regionList.length; i++) {
                listx.push(d.regionList[i].name);
            }

            var option = {
                tooltip: {
                    trigger: 'item',
                    formatter: "{a} <br/>{b} : {c} ({d}%)"
                },
                legend: {
                    bottom: 10,
                    left: 'center',
                    data: listx,
                    textStyle: {
                        color: "#fff"
                    }
                },
                series: [{
                    name: '来源地区',
                    type: 'pie',
                    radius: '60%',
                    center: ['50%', '40%'],
                    label: {
                        normal: {
                            show: false,
                            position: 'center'
                        }
                    },
                    data: d.regionList,
                    itemStyle: {
                        emphasis: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }]
            };


            main1.setOption(option);
        }
    });
}


function main2_data() {
    $.ajax({
        type: "GET",
        url: "/data/get/country",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var listx = [];

            for (var i = 0; i < d.regionList.length; i++) {
                listx.push(d.regionList[i].name);
            }

            var option = {
                tooltip: {
                    trigger: 'item',
                    formatter: "{a} <br/>{b} : {c} ({d}%)"
                },
                legend: {
                    bottom: 10,
                    left: 'center',
                    data: listx,
                    textStyle: {
                        color: "#fff"
                    }
                },
                series: [{
                    name: '来源国家',
                    type: 'pie',
                    radius: '60%',
                    center: ['50%', '40%'],
                    label: {
                        normal: {
                            show: false,
                            position: 'center'
                        }
                    },
                    data: d.regionList,
                    itemStyle: {
                        emphasis: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }]
            };

            main2.setOption(option);
        }
    });

}


function main3_data() {
    $.ajax({
        type: "GET",
        url: "/data/get/ip",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var option = {
                tooltip: {
                    trigger: 'item',
                    formatter: "{a} <br/>{b} : {c} ({d}%)"
                },
                series: [{
                    name: 'IP分布',
                    type: 'pie',
                    radius: '30%',
                    center: ['50%', '50%'],
                    data: d.ipList,
                    itemStyle: {
                        emphasis: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }]
            };

            main3.setOption(option);
        }
    });
}


function main4_data() {
    $.ajax({
        type: "GET",
        url: "/data/get/type",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var dataN = [];
            var dataV = [];

            for (var i = d.typeList.length - 1; i >= 0; i--) {
                dataN.push(d.typeList[i].name);
                dataV.push(d.typeList[i].value);
            }

            var option = {
                tooltip: {},
                xAxis: {
                    type: 'value',
                    show: false
                },
                yAxis: {
                    type: 'category',
                    data: dataN,
                    axisLine: {
                        lineStyle: {
                            color: '#67c8fd'
                        }
                    }
                },
                grid: {
                    left: '0%',
                    right: '0%',
                    bottom: '0%',
                    top: '3%',
                    containLabel: true
                },
                series: [{
                    data: dataV,
                    type: 'bar',
                    label: {
                        normal: {
                            show: true,
                            position: 'inside',
                            color: '#fff'
                        }
                    },
                    itemStyle: {
                        normal: {
                            color: '#49b2c0'
                        }
                    }
                }]
            };

            main4.setOption(option);
        }
    });
}

function main5_data() {
    $.ajax({
        type: "GET",
        url: "/data/get/account",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var option = {

                tooltip: {
                    show: true
                },
                series: [{
                    name: '账号',
                    type: "wordCloud",
                    gridSize: 6,
                    shape: 'circle',
                    sizeRange: [12, 50],
                    width: '100%',
                    height: '80%',
                    textStyle: {
                        normal: {
                            color: function () {
                                return 'rgb(' + [
                                    Math.round(Math.random() * 160),
                                    Math.round(Math.random() * 160),
                                    Math.round(Math.random() * 160)
                                ].join(',') + ')';
                            }
                        },
                        emphasis: {
                            shadowBlur: 10,
                            shadowColor: '#333'
                        }
                    },
                    data: d
                }]

            };

            main5.setOption(option);
        }
    });
}


function main6_data() {
    $.ajax({
        type: "GET",
        url: "/data/get/password",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var option = {

                tooltip: {
                    show: true
                },
                series: [{
                    name: '密码',
                    type: "wordCloud",
                    gridSize: 6,
                    shape: 'circle',
                    sizeRange: [12, 50],
                    width: '100%',
                    height: '80%',
                    textStyle: {
                        normal: {
                            color: function () {
                                return 'rgb(' + [
                                    Math.round(Math.random() * 160),
                                    Math.round(Math.random() * 160),
                                    Math.round(Math.random() * 160)
                                ].join(',') + ')';
                            }
                        },
                        emphasis: {
                            shadowBlur: 10,
                            shadowColor: '#333'
                        }
                    },
                    data: d
                }]

            };

            main6.setOption(option);
        }
    });
}

function mainInfo() {
    $.ajax({
        type: "GET",
        url: "/data/get/info",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var _h = '';

            var result = d.result;
            for (var i = 0; i < result.length; i++) {
                _h += '    <tr class="data_list">';
                _h += '        <td>' + filterXSS(result[i].type) + '</td>';
                _h += '        <td>' + filterXSS(result[i].agent) + '</td>';
                _h += '        <td>' + filterXSS(result[i].ip) + '</td>';
                _h += '        <td>' + filterXSS(result[i].country) + ' ' + filterXSS(result[i].region) + '</td>';
                _h += '        <td>' + filterXSS(formatDate(result[i].create_time)).split(" ")[1] + '</td>';
                _h += '    </tr>';
            }

            $("#info_list").append(_h);
        }
    });
}


function mainWs() {
    var wsuri = "ws://" + window.location.host + "/data/ws";
    var sock = new WebSocket(wsuri);

    sock.onopen = function () {
        console.log("connected to " + wsuri);
    };

    sock.onclose = function (e) {
        console.log("connection closed (" + e.code + ")");
    };

    sock.onmessage = function (e) {
        var data = JSON.parse(e.data);

        if (data.data.type === "center_data") {
            var d = data.data.data;

            var _h = '';

            _h += '    <tr class="data_list">';
            _h += '        <td>' + filterXSS(d.type) + '</td>';
            _h += '        <td>' + filterXSS(d.agent) + '</td>';
            _h += '        <td>' + filterXSS(d.ipx) + '</td>';
            _h += '        <td>' + filterXSS(d.country) + ' ' + filterXSS(d.region) + '</td>';
            _h += '        <td>' + filterXSS(formatDate(d.time)).split(" ")[1] + '</td>';
            _h += '    </tr>';

            $(".data_list").before(_h);
        }
    }
}


function main7_data() {

    function randomData() {
        return Math.round(Math.random() * 300);
    }

    var geoCoordMap = {
        "阿富汗": [67.709953, 33.93911],
        "安哥拉": [17.873887, -11.202692],
        "阿尔巴尼亚": [20.168331, 41.153332],
        "阿联酋": [53.847818, 23.424076],
        "阿根廷": [-63.61667199999999, -38.416097],
        "亚美尼亚": [45.038189, 40.069099],
        "法属南半球和南极领地": [69.348557, -49.280366],
        "澳大利亚": [133.775136, -25.274398],
        "奥地利": [14.550072, 47.516231],
        "阿塞拜疆": [47.576927, 40.143105],
        "布隆迪": [29.918886, -3.373056],
        "比利时": [4.469936, 50.503887],
        "贝宁": [2.315834, 9.30769],
        "布基纳法索": [-1.561593, 12.238333],
        "孟加拉国": [90.356331, 23.684994],
        "保加利亚": [25.48583, 42.733883],
        "巴哈马": [-77.39627999999999, 25.03428],
        "波斯尼亚和黑塞哥维那": [17.679076, 43.915886],
        "白俄罗斯": [27.953389, 53.709807],
        "伯利兹": [-88.49765, 17.189877],
        "百慕大": [-64.7505, 32.3078],
        "玻利维亚": [-63.58865299999999, -16.290154],
        "巴西": [-51.92528, -14.235004],
        "文莱": [114.727669, 4.535277],
        "不丹": [90.433601, 27.514162],
        "博茨瓦纳": [24.684866, -22.328474],
        "中非共和国": [20.939444, 6.611110999999999],
        "加拿大": [-106.346771, 56.130366],
        "瑞士": [8.227511999999999, 46.818188],
        "智利": [-71.542969, -35.675147],
        "中国": [104.195397, 35.86166],
        "象牙海岸": [-5.547079999999999, 7.539988999999999],
        "喀麦隆": [12.354722, 7.369721999999999],
        "刚果民主共和国": [21.758664, -4.038333],
        "刚果共和国": [15.827659, -0.228021],
        "哥伦比亚": [-74.297333, 4.570868],
        "哥斯达黎加": [-83.753428, 9.748916999999999],
        "古巴": [-77.781167, 21.521757],
        "北塞浦路斯": [33.429859, 35.126413],
        "塞浦路斯": [33.429859, 35.126413],
        "捷克共和国": [15.472962, 49.81749199999999],
        "德国": [10.451526, 51.165691],
        "吉布提": [42.590275, 11.825138],
        "丹麦": [9.501785, 56.26392],
        "多明尼加共和国": [-70.162651, 18.735693],
        "阿尔及利亚": [1.659626, 28.033886],
        "厄瓜多尔": [-78.18340599999999, -1.831239],
        "埃及": [30.802498, 26.820553],
        "厄立特里亚": [39.782334, 15.179384],
        "西班牙": [-3.74922, 40.46366700000001],
        "爱沙尼亚": [25.013607, 58.595272],
        "埃塞俄比亚": [40.489673, 9.145000000000001],
        "芬兰": [25.748151, 61.92410999999999],
        "斐": [178.065032, -17.713371],
        "福克兰群岛": [-59.523613, -51.796253],
        "法国": [2.213749, 46.227638],
        "加蓬": [11.609444, -0.803689],
        "英国": [-3.435973, 55.378051],
        "格鲁吉亚": [-82.9000751, 32.1656221],
        "加纳": [-1.023194, 7.946527],
        "几内亚": [-9.696645, 9.945587],
        "冈比亚": [-15.310139, 13.443182],
        "几内亚比绍": [-15.180413, 11.803749],
        "赤道几内亚": [10.267895, 1.650801],
        "希腊": [21.824312, 39.074208],
        "格陵兰": [-42.604303, 71.706936],
        "危地马拉": [-90.23075899999999, 15.783471],
        "法属圭亚那": [-53.125782, 3.933889],
        "圭亚那": [-58.93018, 4.860416],
        "洪都拉斯": [-86.241905, 15.199999],
        "克罗地亚": [15.2, 45.1],
        "海地": [-72.285215, 18.971187],
        "匈牙利": [19.503304, 47.162494],
        "印尼": [113.921327, -0.789275],
        "印度": [78.96288, 20.593684],
        "爱尔兰": [-8.24389, 53.41291],
        "伊朗": [53.688046, 32.427908],
        "伊拉克": [43.679291, 33.223191],
        "冰岛": [-19.020835, 64.963051],
        "以色列": [34.851612, 31.046051],
        "意大利": [12.56738, 41.87194],
        "牙买加": [-77.297508, 18.109581],
        "约旦": [36.238414, 30.585164],
        "日本": [138.252924, 36.204824],
        "哈萨克斯坦": [66.923684, 48.019573],
        "肯尼亚": [37.906193, -0.023559],
        "吉尔吉斯斯坦": [74.766098, 41.20438],
        "柬埔寨": [104.990963, 12.565679],
        "韩国": [127.766922, 35.907757],
        "科索沃": [20.902977, 42.6026359],
        "科威特": [47.481766, 29.31166],
        "老挝": [102.495496, 19.85627],
        "黎巴嫩": [35.862285, 33.854721],
        "利比里亚": [-9.429499000000002, 6.428055],
        "利比亚": [17.228331, 26.3351],
        "斯里兰卡": [80.77179699999999, 7.873053999999999],
        "莱索托": [28.233608, -29.609988],
        "立陶宛": [23.881275, 55.169438],
        "卢森堡": [6.129582999999999, 49.815273],
        "拉脱维亚": [24.603189, 56.879635],
        "摩洛哥": [-7.092619999999999, 31.791702],
        "摩尔多瓦": [28.369885, 47.411631],
        "马达加斯加": [46.869107, -18.766947],
        "墨西哥": [-102.552784, 23.634501],
        "马其顿": [21.745275, 41.608635],
        "马里": [-3.996166, 17.570692],
        "缅甸": [95.956223, 21.913965],
        "黑山": [19.37439, 42.708678],
        "蒙古": [103.846656, 46.862496],
        "莫桑比克": [35.529562, -18.665695],
        "毛里塔尼亚": [-10.940835, 21.00789],
        "马拉维": [34.301525, -13.254308],
        "马来西亚": [101.975766, 4.210484],
        "纳米比亚": [18.49041, -22.95764],
        "新喀里多尼亚": [165.618042, -20.904305],
        "尼日尔": [8.081666, 17.607789],
        "尼日利亚": [8.675277, 9.081999],
        "尼加拉瓜": [-85.207229, 12.865416],
        "荷兰": [5.291265999999999, 52.132633],
        "挪威": [8.468945999999999, 60.47202399999999],
        "尼泊尔": [84.12400799999999, 28.394857],
        "新西兰": [174.885971, -40.900557],
        "阿曼": [55.923255, 21.512583],
        "巴基斯坦": [69.34511599999999, 30.375321],
        "巴拿马": [-80.782127, 8.537981],
        "秘鲁": [-75.015152, -9.189967],
        "菲律宾": [121.774017, 12.879721],
        "巴布亚新几内亚": [143.95555, -6.314992999999999],
        "波兰": [19.145136, 51.919438],
        "波多黎各": [-66.590149, 18.220833],
        "北朝鲜": [127.510093, 40.339852],
        "葡萄牙": [-8.224454, 39.39987199999999],
        "巴拉圭": [-58.443832, -23.442503],
        "卡塔尔": [51.183884, 25.354826],
        "罗马尼亚": [24.96676, 45.943161],
        "俄罗斯": [105.318756, 61.52401],
        "卢旺达": [29.873888, -1.940278],
        "西撒哈拉": [-12.885834, 24.215527],
        "沙特阿拉伯": [45.079162, 23.885942],
        "苏丹": [30.217636, 12.862807],
        "南苏丹": [31.3069788, 6.876991899999999],
        "塞内加尔": [-14.452362, 14.497401],
        "所罗门群岛": [160.156194, -9.64571],
        "塞拉利昂": [-11.779889, 8.460555],
        "萨尔瓦多": [-88.89653, 13.794185],
        "索马里兰": [46.8252838, 9.411743399999999],
        "索马里": [46.199616, 5.152149],
        "塞尔维亚共和国": [21.005859, 44.016521],
        "苏里南": [-56.027783, 3.919305],
        "斯洛伐克": [19.699024, 48.669026],
        "斯洛文尼亚": [14.995463, 46.151241],
        "瑞典": [18.643501, 60.12816100000001],
        "斯威士兰": [31.465866, -26.522503],
        "叙利亚": [38.996815, 34.80207499999999],
        "乍得": [18.732207, 15.454166],
        "多哥": [0.824782, 8.619543],
        "泰国": [100.992541, 15.870032],
        "塔吉克斯坦": [71.276093, 38.861034],
        "土库曼斯坦": [59.556278, 38.969719],
        "东帝汶": [125.727539, -8.874217],
        "特里尼达和多巴哥": [-61.222503, 10.691803],
        "突尼斯": [9.537499, 33.886917],
        "土耳其": [35.243322, 38.963745],
        "坦桑尼亚联合共和国": [34.888822, -6.369028],
        "乌干达": [32.290275, 1.373333],
        "乌克兰": [31.16558, 48.379433],
        "乌拉圭": [-55.765835, -32.522779],
        "美国": [-95.712891, 37.09024],
        "乌兹别克斯坦": [64.585262, 41.377491],
        "委内瑞拉": [-66.58973, 6.42375],
        "越南": [108.277199, 14.058324],
        "瓦努阿图": [166.959158, -15.376706],
        "西岸": [35.3027226, 31.9465703],
        "也门": [48.516388, 15.552727],
        "南非": [22.937506, -30.559482],
        "赞比亚": [27.849332, -13.133897],
        "津巴布韦": [29.154857, -19.015438],
        '新疆': [87.500966, 43.983832],
        '西藏': [90.959657, 29.881987],
        '青海': [101.703679, 36.733408],
        '甘肃': [103.764176, 36.198433],
        '内蒙古': [111.711808, 40.98898],
        '宁夏': [106.192619, 38.605171],
        '四川': [103.984944, 30.712171],
        '云南': [102.733927, 25.025991],
        '陕西': [112.521289, 38.025365],
        '山西': [108.84183, 34.510421],
        '重庆': [106.413387, 29.689402],
        '贵州': [106.560565, 26.756654],
        '广西': [108.326706, 22.99805],
        '海南': [110.129641, 20.14162],
        '广东': [113.183592, 23.202287],
        '澳门': [113.551538, 22.109432],
        '香港': [114.066662, 22.588638],
        '台湾': [121.49917, 25.12653],
        '福建': [119.107522, 26.193691],
        '江西': [115.722419, 28.882959],
        '湖南': [112.778851, 28.363482],
        '湖北': [114.177046, 30.743959],
        '安徽': [117.120614, 31.943998],
        '浙江': [119.990592, 30.361806],
        '江苏': [118.665986, 32.194658],
        '河南': [113.441154, 34.8448],
        '山东': [116.973435, 36.763019],
        '上海': [121.315197, 31.314325],
        '河北': [114.397814, 38.170754],
        '天津': [117.194203, 39.180291],
        '北京': [116.384722, 39.977552],
        '辽宁': [123.412489, 41.875105],
        '吉林': [125.252219, 43.850841],
        '黑龙江': [126.503235, 45.865719]
    };

    var convertData = function (data) {
        var res = [];
        for (var i = 0; i < data.length; i++) {
            var dataItem = data[i];
            var fromCoord = geoCoordMap[dataItem[0].name];
            var toCoord = geoCoordMap[dataItem[1].name];
            if (fromCoord && toCoord) {
                res.push([{
                    coord: fromCoord,
                    value: dataItem[0].value
                },
                    {
                        coord: toCoord
                    }
                ]);
            }
        }
        return res;
    };


    $.ajax({
        type: "GET",
        url: "/data/get/word",
        dataType: "json",
        success: function (e) {
            var d = e.data;

            var geoCoordArry = [];
            for (var item in geoCoordMap) {
                geoCoordArry.push(item);
            }

            var attckData = [];
            for (var i = 0; i < d.length; i++) {

                if (geoCoordArry.indexOf(d[i].name) > -1) {
                    var acctckJson = [
                        {
                            "name": d[i].name,
                            "value": d[i].value
                        },
                        {
                            "name": attckCity
                        }
                    ];

                    attckData.push(acctckJson)
                }
            }

            var series = [];
            [
                [attckCity, attckData]
            ].forEach(function (item, i) {
                series.push({
                        name: '攻击线1',
                        type: "lines",
                        zlevel: 2,
                        effect: {
                            show: true,
                            color: "#0bc7f3",
                            period: 4, //箭头指向速度，值越小速度越快
                            trailLength: 0.02, //特效尾迹长度[0,1]值越大，尾迹越长重
                            symbol: "arrow", //箭头图标
                            symbolSize: 5 //图标大小
                        },
                        lineStyle: {
                            normal: {
                                color: '#0bc7f3',
                                width: 1, //尾迹线条宽度
                                opacity: 0, //尾迹线条透明度
                                curveness: 0.3 //尾迹线条曲直度
                            }
                        },
                        data: convertData(item[1])
                    }, {
                        type: "effectScatter",
                        coordinateSystem: "geo",
                        zlevel: 2,
                        rippleEffect: {
                            //涟漪特效
                            period: 4, //动画时间，值越小速度越快
                            brushType: "stroke", //波纹绘制方式 stroke, fill
                            scale: 4 //波纹圆环最大限制，值越大波纹越大
                        },
                        symbol: "circle",
                        symbolSize: function (val) {
                            return 4 + val[2] / 1000; //圆环大小
                        },
                        itemStyle: {
                            normal: {
                                show: true,
                            },
                            emphasis: {
                                show: true,
                                color: "#FF6A6A"
                            }
                        },
                        data: item[1].map(function (dataItem) {
                            return {
                                name: dataItem[0].name,
                                value: geoCoordMap[dataItem[0].name].concat([dataItem[0].value])
                            };
                        })
                    },
                    //被攻击点
                    {
                        type: "scatter",
                        coordinateSystem: "geo",
                        zlevel: 2,
                        rippleEffect: {
                            period: 4,
                            brushType: "stroke",
                            scale: 4
                        },
                        label: {
                            normal: {
                                show: true,
                                color: "red",
                                position: "right",
                                formatter: "{b}",
                            },
                            emphasis: {
                                show: true,
                                color: "#FF6A6A"
                            }
                        },
                        symbol: "pin",
                        symbolSize: 25,
                        itemStyle: {
                            normal: {
                                show: true,
                                color: "red",
                            },
                            emphasis: {
                                show: true,
                                color: "#FF6A6A"
                            }
                        },
                        data: [{
                            name: item[0],
                            value: geoCoordMap[item[0]].concat([100]),
                            visualMap: false
                        }]
                    }
                );
            });

            var worldMapOpt = {
                tooltip: {
                    trigger: "item",
                    backgroundColor: "#1540a1",
                    borderColor: "#FFFFCC",
                    showDelay: 0,
                    hideDelay: 0,
                    enterable: true,
                    transitionDuration: 0,
                    extraCssText: "z-index:100",
                    formatter: function (params, ticket, callback) {
                        var res = "";
                        var name = params.name;
                        var value = typeof(params.value[params.seriesIndex + 1]) == "undefined" ? 0 : params.value[params.seriesIndex + 1];
                        res =
                            "<span style='color:#fff;'>" +
                            name +
                            "</span><br/>数据：" +
                            value;
                        return res;
                    }
                },
                visualMap: {
                    //图例值控制
                    show: false,
                    type: 'piecewise',
                    pieces: [{
                        max: 80,
                        color: '#00e200'
                    },
                        {
                            min: 80,
                            max: 120,
                            color: '#f5ff00'
                        },
                        {
                            min: 120,
                            color: '#ff0500'
                        }
                    ],
                    calculable: true
                },
                geo: {
                    map: "world",
                    show: true,
                    label: {
                        emphasis: {
                            show: false
                        }
                    },
                    roam: true,
                    layoutCenter: ["50%", "50%"], //地图位置
                    layoutSize: "190%",
                    itemStyle: {
                        normal: {
                            show: 'true',
                            color: "#04284e", //地图背景色
                            borderColor: "#5bc1c9" //省市边界线
                        },
                        emphasis: {
                            show: 'true',
                            color: "rgba(37, 43, 61, .5)" //悬浮背景
                        }
                    }
                },
                legend: {
                    orient: 'vertical',
                    top: '30',
                    left: 'center',
                    align: 'right',
                    textStyle: {
                        color: '#fff',
                        fontSize: 20
                    },
                    itemWidth: 50,
                    itemHeight: 30,
                    selectedMode: 'multiple'
                },
                series: series
            };

            worldMap.setOption(worldMapOpt);

        }
    });
}

function init_chart() {
    main1_data();
    main2_data();
    main3_data();
    main4_data();
    main5_data();
    main6_data();
    main7_data();
}

init_chart();
mainInfo();
mainWs();

setInterval(init_chart, 60000);

// ==================================

function getTime() {
    var date = new Date();
    var year = date.getFullYear();        //年 ,从 Date 对象以四位数字返回年份
    var month = date.getMonth() + 1;      //月 ,从 Date 对象返回月份 (0 ~ 11) ,date.getMonth()比实际月份少 1 个月
    var day = date.getDate();             //日 ,从 Date 对象返回一个月中的某一天 (1 ~ 31)
    var hours = date.getHours();          //小时 ,返回 Date 对象的小时 (0 ~ 23)
    var minutes = date.getMinutes();      //分钟 ,返回 Date 对象的分钟 (0 ~ 59)
    var seconds = date.getSeconds();      //秒 ,返回 Date 对象的秒数 (0 ~ 59)
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (day >= 0 && day <= 9) {
        day = "0" + day;
    }
    if (hours >= 0 && hours <= 9) {
        hours = "0" + hours;
    }
    if (minutes >= 0 && minutes <= 9) {
        minutes = "0" + minutes;
    }
    if (seconds >= 0 && seconds <= 9) {
        seconds = "0" + seconds;
    }
    var currentFormatDate = year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;

    $("#timex").text(currentFormatDate);
}

function formatDate(d) {
    var datex = new Date(d).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '');
    return datex
}

setInterval(getTime, 1000);

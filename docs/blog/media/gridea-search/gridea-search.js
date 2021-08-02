/**
 * 由 宝硕(https://github.com/renbaoshuo) 修改自 https://github.com/tangkaichuan/gridea-search
 */

(function () {
    //获取url参数
    function getParam(url, param) {
        if (url.indexOf('?') > -1) {
            var urlSearch = url.split('?');
            var paramList = urlSearch[1].split('&');
            for (var i = paramList.length - 1; i >= 0; i--) {
                var temp = paramList[i].split('=');
                if (temp[0] === param) {
                    return temp[1];
                }
            }
        } else {
            return "";
        }
    }

    //获取解码后的搜索词
    function getQueryPhrase() {
        if(window.location.href.indexOf('?') == -1) return "";
        var phrase = getParam(window.location.href, 'q');
        var queryPhrase = decodeURIComponent(phrase.replace(/\+/g, ' '));
        return queryPhrase;
    }

    //填充搜索输入框
    function fillSearchInput() {
        var searchForm = document.getElementById('gridea-search-form');
        var searchInput = searchForm.getElementsByTagName('input')[0];
        searchInput.value = getQueryPhrase();
    }

    //Ajax异步GET请求
    function ajax(obj) {
        var xhr = new XMLHttpRequest();
        xhr.open('get', obj.url, true);
        xhr.send(null);
        xhr.onreadystatechange = function () {
            //异步请求：响应状态为4，数据加载完毕
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    obj.success(xhr.responseText);
                } else {
                    obj.error(xhr.status);
                }
            }
        }
    }

    //模糊搜索 https://github.com/krisk/fuse
    function fuzzySearch(data, phrase) {
        var options = {
            includeMatches: true,
            ignoreLocation: true,
            keys: [
                'title',
                'content'
            ]
        };
        var fuse = new Fuse(data, options);
        var fuzzyResult = fuse.search(phrase);
        return fuzzyResult;
    } 

    //检查缓存是否最新
    function checkCache() {
        var caches = {};
        caches.infos = JSON.parse(localStorage.getItem('InfosCache'));
        caches.contents = JSON.parse(localStorage.getItem('ContentsCache'));
        if (caches.infos && caches.contents) {
            var cachedTime = caches.infos.utils.now.toString();
            var updateTime = document.getElementById('gridea-search-result').getAttribute('data-update');
            if (cachedTime === updateTime) {
                return caches;
            }
        }
        localStorage.removeItem('InfosCache');
        localStorage.removeItem('ContentsCache');
        return false;
    }

    //获取博客全文api
    function getContents(callback) {
        var caches = checkCache();
        if (caches) {
            callback(caches.contents);
        }
        else {
            ajax({
                url: '../api-content/index.html' + "?_=" + Date.now(),
                success: function (data) {
                    callback(JSON.parse(data));
                    localStorage.setItem('ContentsCache', data);
                }
            });
        }
    }

    //获取博客信息api
    function getInfos(callback) {
        var caches = checkCache();
        if (caches) {
            callback(caches.infos);
        }
        else {
            ajax({
                url: '../api-info/index.html' + "?_=" + Date.now(),
                success: function (data) {
                    callback(JSON.parse(data));
                    localStorage.setItem('InfosCache', data);
                }
            });
        }
    }

    //根据一段文本调用模糊搜索
    function searchBy(phrase, callback) {
        var result = '';
        var getFuzzyResult = function (data) {
            result = fuzzySearch(data.posts, phrase);
            callback(result);
        }
        //根据全文内容获取搜索结果
        getContents(getFuzzyResult);
    }

    //显示无搜索结果
    function showNoResult() {
        var resultDIV = document.getElementById('gridea-search-result');
        var noResult = resultDIV.getElementsByClassName('no-result')[0];
        noResult.style.display = 'block';
        resultDIV.innerHTML = noResult.outerHTML;
    }

    //根据解码后的搜索词执行搜索
    function searchByPhrase(resultHandler) {
        var queryPhrase = getQueryPhrase();
        if (queryPhrase === '' || typeof (queryPhrase) === 'undefined') {
            showNoResult();
        } else {
            searchBy(queryPhrase, resultHandler);
        }
    }

    //获取搜索结果列表模板的URL
    function getTemplateURL() {
        var scripts = document.getElementsByTagName('script');
        var templateURL = '';
        for (var i = 0; i < scripts.length; i++) {
            if (scripts[i].type === 'text/ejs') {
                templateURL = scripts[i].src;
                return templateURL;
            }
        }
    }

    //渲染搜索结果列表ejs https://github.com/mde/ejs
    function renderResult(searchedInfos) {
        if (searchedInfos.posts.length > 0) {
            ajax({
                url: getTemplateURL(),
                success: function (data) {
                    var resultDIV = document.getElementById('gridea-search-result');
                    resultDIV.innerHTML = ejs.compile(data)(searchedInfos);
                }
            });
        } else {
            showNoResult();
        }
    }

    //搜索结果关键字高亮
    function keywordHighlight(searchedContent) {
        var searchedPostContent = searchedContent.item.content;//搜索结果内容预览
        var preview = '... ';
        for (var i = 0; i < searchedContent.matches.length; i++) {
            if (searchedContent.matches[i].key === 'content') {//如果匹配到文章内容，截取关键字
                var indices = searchedContent.matches[i].indices[0];
                var beforeKeyword = searchedPostContent.substring(indices[0] - 20, indices[0]);//关键字前10字
                var keyword = searchedPostContent.substring(indices[0], indices[1] + 1);//关键字
                var afterKeyword = searchedPostContent.substring(indices[1] + 1, indices[1] + 120);//关键字后70字
                preview = beforeKeyword + '<span class="searched-keyword">'
                    + keyword + '</span>' + afterKeyword;
            } else {//没有匹配到文章内容，则是标题，直接截取前80字
                preview = searchedPostContent.substring(0, 80);
            }
        }
        return preview + ' ...';
    }

    //循环匹配搜索到的内容与展示信息
    function getResult(infos, searchedContents) {
        var searchedInfos = JSON.parse(JSON.stringify(infos));//对象深拷贝
        searchedInfos.posts = [];
        for (var i = 0; i < searchedContents.length; i++) {
            for (var j = 0; j < infos.posts.length; j++) {
                if (searchedContents[i].item.link === infos.posts[j].link) {
                    infos.posts[j].searchedPreview = keywordHighlight(searchedContents[i]);// 预览关键字高亮
                    infos.posts[j].content = searchedContents[i].item.content;// content注入
                    searchedInfos.posts.push(infos.posts[j]);// push到所需结果中
                }
            }
        }
        return searchedInfos;
    }

    //主方法
    function grideaSearch() {
        //搜索结果回调
        var resultHandler = function (searchedContents) {
            getInfos(function (infos) {
                // console.log(infos);
                // console.log(searchedContents);
                var searchedInfos = getResult(infos, searchedContents);
                renderResult(searchedInfos);
            });
        }
        searchByPhrase(resultHandler);
    }

    //页面加载完执行
    window.onload = function () {
        fillSearchInput();
        grideaSearch();
    }
})();

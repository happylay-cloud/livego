# 访问地址
```
http://127.0.0.1:8090/statics/
```
# 兼容性报告-已测试

搜狗浏览器暂不兼容

华为荣耀v9浏览器暂不兼容、微信不兼容

小米兼容

# 测试时需要修改index.html中src地址

```http request
video.js 

7.7.5（最高兼容版本-不兼容safari）

7.5.5 

7.4.1（最低兼容版本-兼容safari）
```
```scss
/* 显示时间轴-bug */
.video-js .vjs-time-control {
    display: block;
}

.video-js .vjs-remaining-time {
    display: none;
}

/* 进度条 */
video::-webkit-media-controls-timeline {
    display: none;
}

video {
    background-color: #000;
}

.vjs-big-play-button {
    left: 50% !important;
    top: 50% !important;
    margin-top: -0.5rem!important;
    margin-left: -0.5rem!important;
    border-radius: 2rem!important;
    width: 1rem!important;
    height: 1rem!important;
    line-height: 1rem!important;
    border-color: #0B5ED7!important;
    background-color: #0B5ED7!important;
}

.video-js .vjs-time-tooltip{
    display: none!important;
}

```
```scss
<style>
    /* 给.video-js设置字体大小以统一各浏览器样式表现，因为video.js采用的是em单位 */
    .video-js{
        font-size: 14px;
    }

    .video-js button{
        outline: none;
    }

    /* 视频占满容器高度 */
    .video-js.vjs-fluid,
    .video-js.vjs-16-9,
    .video-js.vjs-4-3{
        height: 100%;
        background-color: #161616;
    }
    .vjs-poster{
        background-color: #161616;
    }
    /* 中间大的播放按钮 */
    .video-js .vjs-big-play-button{
        font-size: 2.5em;
        line-height: 2.3em;
        height: 2.5em;
        width: 2.5em;
        -webkit-border-radius: 2.5em;
        -moz-border-radius: 2.5em;
        border-radius: 2.5em;
        background-color: rgba(115,133,159,.5);
        border-width: 0.12em;
        margin-top: -1.25em;
        margin-left: -1.75em;
    }

    /* 视频暂停时显示播放按钮 */
    .video-js.vjs-paused .vjs-big-play-button{
        display: block;
    }

    /* 视频加载出错时隐藏播放按钮 */
    .video-js.vjs-error .vjs-big-play-button{
        display: none;
    }

    /* 加载圆圈 */
    .vjs-loading-spinner {
        font-size: 2.5em;
        width: 2em;
        height: 2em;
        border-radius: 1em;
        margin-top: -1em;
        margin-left: -1.5em;
    }

    /* 控制条默认显示 */
    .video-js .vjs-control-bar{
        display: flex;
    }

    /* 显示时间轴-bug */
    .video-js .vjs-time-control{
        display:block;
    }
    .video-js .vjs-remaining-time{
        display: none;
    }

    /* 控制条所有图标，图标字体大小最好使用px单位，如果使用em，各浏览器表现可能会不大一样 */
    .vjs-button > .vjs-icon-placeholder:before{
        font-size: 22px;
        line-height: 1.9;
    }
    .video-js .vjs-playback-rate .vjs-playback-rate-value{
        line-height: 2.4;
        font-size: 18px;
    }
    
    /* 进度条背景色 */
    .video-js .vjs-play-progress{
        color: #2A82D9;
        background-color: #2A82D9;
    }

    /* 鼠标悬浮提示背景色 */
    .video-js .vjs-progress-control .vjs-mouse-display{
        background-color: #2A82D9;
    }

    /* 不显示鼠标悬浮提示-兼容性考虑 */
    .video-js .vjs-time-tooltip{
        display: none!important;
    }

</style>
```

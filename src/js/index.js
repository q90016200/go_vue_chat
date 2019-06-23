import head from './head';
import Vue from 'vue';

head();

var app = new Vue({
    el: '#app',
    mounted: function () {
        console.log('Hello Webpack and Vue !');
    },
    data: {
        message: 'Hello Vue!'
    }
});

app.message = 'Hello Vue!!';
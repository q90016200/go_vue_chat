const path = require('path');

module.exports = {
    entry: './src/js/index.js',
    mode:'development',
    // mode:'production',
    resolve: {
        alias: {
            vue: 'vue/dist/vue.esm.js'
        }
    },
    output: {
        path: path.resolve(__dirname, 'public'),
        filename: 'index.js'
    }
};
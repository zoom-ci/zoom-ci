module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:7002/',
                changeOrigin: true,
            }
        }
    },
    publicPath: '/web/dist',
    outputDir: "dist"
};

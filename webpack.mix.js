const mix = require('laravel-mix');

const res_base_path = "./src/resources";
const dist_base_path = "./src/public/assets";

mix.disableNotifications();
mix.setPublicPath("./src/public");

mix.js(`${res_base_path}/js/index.js`, `${dist_base_path}/js/`);
mix.postCss(`${res_base_path}/css/index.css`, `${dist_base_path}/css/`);
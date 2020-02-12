const guid2 = ()=>{
    function S4() {
        return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
    }
    return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}
const Utils = {
    guid2
}
const install = (Vue) => {
    Vue.prototype.Utils = Utils
}

export default install
export {Utils}
class List {
    constructor() {
        this.list = [];
    }

    /**
     * 添加一个节点
     * @param item 
     */
    setNode(item) {
        this.list.push(item);
    }

    /**
     * 获取一个节点
     */
    getNode() {
        return this.list.shift();
    }

    /**
     * 返回所有的节点
     */
    getList() {
        return this.list;
    }
}

export default List;
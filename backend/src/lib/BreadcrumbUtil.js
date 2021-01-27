import List from "@/lib/datastructure/List"

class BreadcrumbUtil {
    constructor() {
        this.list = new List();
    }

    /**
     * 添加面包屑
     * @param name 
     * @param href 
     */
    setItem(name, href) {
        let bread = new Bread(name, href);
        this.list.setNode(bread);
    }

    /**
     * 获取面包屑
     */
    getBreadcrumbs() {
        return this.list.getList();
    }
}

class Bread {
    constructor(name, href) {
        this.name = name;
        this.href = href;
    }

    getBread() {
        return {
            'name': this.name,
            'href': this.href
        };
    }
}


export {
    BreadcrumbUtil,
    Bread
};
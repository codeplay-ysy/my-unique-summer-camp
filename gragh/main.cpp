#include "C:\myproject\summercamp\my-unique-summer-camp\gragh\Gragh.h"  // 包含Graph类的头文件

int main() {
    Graph graph;  // 创建Graph对象

    // 使用邻接矩阵法存储图
    graph.AM_GraphInitial();
    graph.AM_GraphPrint();
    graph.AM_BFS(0); // 进行广度优先搜索
    graph.AM_DFS(0); // 进行深度优先搜索

    // 使用邻接表法存储图
    graph.AL_GraphInitial();
    graph.AL_GraphPrint();
    graph.AL_BFS(0); // 进行广度优先搜索
    graph.AL_DFS(0); // 进行深度优先搜索

    return 0;
}

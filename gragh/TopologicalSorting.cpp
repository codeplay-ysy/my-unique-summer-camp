#include <iostream>
#include <stack>
#define MAX_VERTEX_NUM 20
#define OK 1
#define ERROR 0
using namespace std;
typedef char NumType;
typedef int Status;

//下面是邻接表构造有向图过程
struct ArcNode
{
    int AdjVex;
    ArcNode *NextArc;
};
struct VexNode
{
    NumType data;
    int InDegree;
    ArcNode *FirstArc;
};
struct ALGraph
{
    VexNode Vex[MAX_VERTEX_NUM];
    int VexNum;
    int ArcNum;
};
int Locate(ALGraph G,NumType v)
{
    int i;
    for(i=0;i<G.VexNum;i++)
        if(v==G.Vex[i].data)return i;
    return -1;
}
void CreatALGraph(ALGraph &G)
{
    cout<<"请输入顶点数和弧数："<<endl;
    cin>>G.VexNum;cin>>G.ArcNum;
    int i;
    cout<<"请输入顶点数据："<<endl;
    for(i=0;i<G.VexNum;i++)
    {
        cin>>G.Vex[i].data;
        G.Vex[i].FirstArc=0;
        G.Vex[i].InDegree=0;
    }
    NumType v1,v2;
    int j,k;
    cout<<"请输入弧："<<endl;
    for(k=0;k<G.ArcNum;k++)
    {
        cin>>v1;cin>>v2;
        i=Locate(G,v1);j=Locate(G,v2);
        G.Vex[j].InDegree++;
        ArcNode *p=new ArcNode();
        *p={j,G.Vex[i].FirstArc};
        G.Vex[i].FirstArc=p;
    }
}

//上面是用邻接表构造有向图
//下面是拓扑排序代码

void TopoLogicalSort(ALGraph G)
{
    ArcNode *p=0;
    stack<int> s;
    int i;
    for(i=0;i<G.VexNum;i++)
        if(G.Vex[i].InDegree==0)s.push(i);
    int t;
    int count0=0;
    while(!s.empty())
    {
        count0++;
        t=s.top();
        s.pop();
        cout<<G.Vex[t].data<<" ";
        for(p=G.Vex[t].FirstArc;p!=0;p=p->NextArc)
        {
            G.Vex[p->AdjVex].InDegree--;
            if(!G.Vex[p->AdjVex].InDegree)s.push(p->AdjVex);
        }
    }

    if(count0==G.VexNum){cout<<"YES";return;}
    cout<<"NO";
}
int main()
{
    ALGraph G;
    CreatALGraph(G);
    TopoLogicalSort(G);
    return 0;
}


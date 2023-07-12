//迭代创建二叉树。。。发现自己之前不会创建，所以试试。
#include <iostream>
using namespace std;
#include <queue>
struct BTreeNode
{
	char data;
	BTreeNode* left;
	BTreeNode* right;
};

class BTree
{
public:
	void levelCreate(BTreeNode*& Node)
	{
		queue<BTreeNode*> que;//创建一个队列,用于存放节点,先进先出
		char data;
		cin >> data;
		if (data != '0')
		{
			Node = new BTreeNode;
			Node->data = data;
			que.push(Node);
		}
		else
		{
			Node = NULL;
			return;
		}
		while (!que.empty())//队列不为空,说明还有节点没有创建
		{
			BTreeNode* node = que.front();
			que.pop();//取出队列中的第一个节点
			//输入左值
			cin >> data;
			if (data != '0')
			{
				node->left = new BTreeNode;
				node->left->data = data;
				que.push(node->left);
			}
			else
			{
				node->left = NULL;
			}
			//输入右值
			cin >> data;
			if (data != '0')
			{
				node->right = new BTreeNode;
				node->right->data = data;
				que.push(node->right);
			}
			else
			{
				node->right = NULL;
			}
		}
	}

	void clear(BTreeNode*& Node)
	{
		if (Node)
		{
			clear(Node->left);
			clear(Node->right);
			delete Node;
		}
	}
};

int main()
{
	BTree tree;
	BTreeNode* boot;
	tree.levelCreate(boot);
	cout << "二叉树创建完成！" << endl;
	system("pause");
	tree.clear(boot);
	cout << "二叉树清理完成！" << endl;
	system("pause");
	return 0;
}

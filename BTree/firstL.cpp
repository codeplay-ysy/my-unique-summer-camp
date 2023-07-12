//迭代实现先序遍历
/*
方案1：
准备一个栈容器，只要遇到一个非空节点，就将其推入栈中，一直向左访问,
直到访问到空指针，然后从栈中弹出元素，依次回溯访问之前还没有访问的
右节点，直到此时的指针为空并且栈中元素为空，说明所有非空节点右边子
节点均以访问，循环停止
方案2：
准备一个栈容器， 将它的右子节点存入栈中， 递推左子节点直到为空，从
栈中弹出一个元素，访问它，一直这样循环直到栈为空，说明所有右子节点
都访问完，即前序遍历完成
*/
#include <iostream>
using namespace std;
#include <stack>
struct BTreeNode
{
	char data;
	BTreeNode* left;
	BTreeNode* right;
};

class BTree
{
public:
	void create(BTreeNode*& Node)
	{
		char data;
		cin >> data;
		if (data != '0')
		{
			Node = new BTreeNode;
			Node->data = data;
			create(Node->left);
			create(Node->right);
		}
		else
		{
			Node = NULL;
		}
	}

	void preorderTree1(BTreeNode* Node)
	{
		stack<BTreeNode*> node;//准备一个栈容器
		BTreeNode* pnode = Node;//pnode代表当前节点
		while (pnode != NULL || !node.empty())
		{
			if (pnode)
			{
				cout << pnode->data << " ";
				node.push(pnode);
				pnode = pnode->left;
			}
			else
			{
				BTreeNode* treenode = node.top();
				node.pop();
				pnode = treenode->right;
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
    void preorderTree2(BTreeNode* Node)//方案2
	{
		stack<BTreeNode*> node;
		while (true)
		{
			while (Node)
			{
				cout << Node->data << " ";
				node.push(Node->right);
				Node = Node->left;
			}
			if (node.empty())
			{
				break;
			}
			Node = node.top();
			node.pop();
		}
	}
};

int main()
{
	BTree tree;
	BTreeNode* root;
	tree.create(root);
	cout << "先序遍历：" << endl;
	tree.preorderTree1(root);//使用了方案1
	cout << endl;
	tree.clear(root);
	system("pause");
	return 0;
}

#include <iostream>
#include <map>
#include <vector>

template<typename Key, typename Value>
class HashTable {
private:
    std::map<Key, std::vector<Value>> hashTable;

public:
    void insert(const Key& key, const Value& value) {
        hashTable[key].push_back(value);
    }

    void remove(const Key& key, const Value& value) {
        auto it = hashTable.find(key);
        if (it != hashTable.end()) {
            auto& values = it->second;
            values.erase(std::remove(values.begin(), values.end(), value), values.end());
            if (values.empty()) {
                hashTable.erase(it);
            }
        }
    }

    void update(const Key& key, const Value& oldValue, const Value& newValue) {
        auto it = hashTable.find(key);
        if (it != hashTable.end()) {
            auto& values = it->second;
            for (auto& val : values) {
                if (val == oldValue) {
                    val = newValue;
                }
            }
        }
    }

    bool find(const Key& key, std::vector<Value>& result) const {
        auto it = hashTable.find(key);
        if (it != hashTable.end()) {
            result = it->second;
            return true;
        }
        return false;
    }

    void printHashTable() const {
        for (const auto& entry : hashTable) {
            std::cout << "Key: " << entry.first << ", Values: ";
            for (const auto& value : entry.second) {
                std::cout << value << " ";
            }
            std::cout << std::endl;
        }
    }
};

int main() {
    HashTable<std::string, int> hashTable;

    hashTable.insert("apple", 10);
    hashTable.insert("banana", 5);
    hashTable.insert("orange", 7);
    hashTable.insert("apple", 15); //测试key相同时的情况

    std::vector<int> values;
    if (hashTable.find("apple", values)) {
        std::cout << "Values for key 'apple': ";
        for (const auto& value : values) {
            std::cout << value << " ";
        }
        std::cout << std::endl;
    }

    hashTable.update("apple", 10, 20);
    hashTable.printHashTable();

    hashTable.remove("apple", 20);
    hashTable.printHashTable();

    return 0;
}

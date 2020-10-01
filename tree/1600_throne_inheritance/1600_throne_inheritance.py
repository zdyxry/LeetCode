from typing import List



class ThroneInheritance:

    def __init__(self, kingName: str):
        self.relations = {kingName: []}
        self.death_order = set()
        self.kind_name = kingName

    def birth(self, parentName: str, childName: str) -> None:
        if self.relations.get(parentName):
            self.relations[parentName] += [childName]
        else:
            self.relations[parentName] = [childName]

    def death(self, name: str) -> None:
        self.death_order.add(name)

    def getInheritanceOrder(self) -> List[str]:
        def dfs(parentName, cur_order):
            if parentName not in self.death_order:
                cur_order.append(parentName)
            if self.relations.get(parentName) is not None:
                for v in self.relations.get(parentName):
                    dfs(v, cur_order)

        cur_order = []            
        dfs(self.kind_name, cur_order)
        return cur_order

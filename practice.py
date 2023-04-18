class Account:
    def __init__(self, account_id, initial_balance=0):
        self.account_id = account_id
        self.balance = initial_balance
        self.transactions = []
        self.total_outgoing = 0

    def deposit(self, amount):
        self.balance += amount
        self.transactions.append(('deposit', amount))

    def withdraw(self, amount):
        if self.balance >= amount:
            self.balance -= amount
            self.transactions.append(('withdraw', amount))
            self.total_outgoing += amount

class BankingSystem:
    def __init__(self):
        self.accounts = {}


    def top_spenders(self, timestamp, n):
        top_accounts = sorted(self.accounts.values(), key=lambda acc: (-acc.total_outgoing, acc.account_id))
        top_n = top_accounts[:n]
        return ', '.join([f"{acc.account_id}({acc.total_outgoing})" for acc in top_n])

    def create_account(self, timestamp, account_id):
        if account_id in self.accounts:
            return "false"
        account = Account(account_id)
        self.accounts[account_id] = account
        return "true"

    def deposit(self, timestamp, account_id, amount):
        if account_id not in self.accounts:
            return ""
        account = self.accounts[account_id]
        account.deposit(int(amount))
        return str(account.balance)

    def transfer(self, timestamp, source_account_id, target_account_id, amount):
        if source_account_id not in self.accounts or target_account_id not in self.accounts:
            return ""
        if source_account_id == target_account_id:
            return ""

        source_account = self.accounts[source_account_id]
        target_account = self.accounts[target_account_id]

        if source_account.balance < int(amount):
            return ""

        source_account.withdraw(int(amount))
        target_account.deposit(int(amount))

        return str(source_account.balance)

    # ... (rest of the class)


    def solution(self, queries):
        results = []

        for query in queries:
            action = query[0].upper()

            if action == "CREATE_ACCOUNT":
                timestamp, account_id = query[1], query[2]
                result = self.create_account(timestamp, account_id)
                results.append(result)

            elif action == "DEPOSIT":
                timestamp, account_id, amount = query[1], query[2], query[3]
                result = self.deposit(timestamp, account_id, amount)
                results.append(result)

            elif action == "TRANSFER":
                timestamp, source_account_id, target_account_id, amount = query[1], query[2], query[3], query[4]
                result = self.transfer(timestamp, source_account_id, target_account_id, amount)
                results.append(result)
            elif action == "TOP_SPENDERS":
                timestamp,n = query[1],int(query[2])
                result = self.top_spenders(timestamp,n)
                results.append(result)

        return results

def solution(queries):
    banking_system = BankingSystem()
    return banking_system.solution(queries)




def main():
    queries = [
        ["CREATE_ACCOUNT", "1", "account1"],
        ["CREATE_ACCOUNT", "2", "account2"],
        ["CREATE_ACCOUNT", "3", "account3"],
        ["DEPOSIT", "4", "account1", "1000"],
        ["DEPOSIT", "5", "account2", "1000"],
        ["DEPOSIT", "6", "account3", "1000"],
        ["TRANSFER", "7", "account2", "account3", "100"],
        ["TRANSFER", "8", "account2", "account1", "100"],
        ["TRANSFER", "9", "account3", "account1", "100"],
        ["TOP_SPENDERS", "10", "3"]
    ]

    results = solution(queries)
    print(results)  # Output: ["true", "true", "true", "1000", "1000", "1000", "900", "800", "1000", "account2(200), account3(100), account1(0)"]

if __name__ == "__main__":
    main()
   
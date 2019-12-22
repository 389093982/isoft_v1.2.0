package test

import "fmt"

type Account struct {
	accountNumber string
	password      string
	balance       float64
}

func (account *Account) Deposit(pwd string, amount float64) {
	if pwd != account.password {
		fmt.Printf("您输入的密码有误！\n")
		return
	}
	if amount <= 0 {
		fmt.Printf("存款的金额不正确！\n")
		return
	}
	account.balance += amount
	fmt.Printf("存款%v成功，目前余额：%v\n", amount, account.balance)
}

func (account *Account) WithDrown(pwd string, amount float64) {
	if pwd != account.password {
		fmt.Printf("您输入的密码有误！\n")
		return
	}
	if amount <= 0 || amount > account.balance {
		fmt.Printf("取款的金额不正确！\n")
		return
	}
	account.balance -= amount
	fmt.Printf("取款%v成功，目前余额：%v\n", amount, account.balance)
}

func (account *Account) Query(pwd string) {
	if pwd != account.password {
		fmt.Printf("您输入的密码有误！\n")
		return
	}
	fmt.Printf("您的账户%v，目前余额：%v\n", account.accountNumber, account.balance)
}

func TestAccount() {
	account := Account{
		accountNumber: "622168",
		password:      "888888",
		balance:       1000,
	}

	for true {
		var operator int
		var password string
		var amount float64
		fmt.Println("=====招商银行欢迎您，您可以办理以下业务=====")
		fmt.Printf("1:存款\n2:取款\n3:查询\n")
		fmt.Print("请选择:")
		fmt.Scanln(&operator)
		switch operator {
		case 1:
			fmt.Print("您正在办理的业务是存款，请输入密码：")
			fmt.Scanln(&password)
			fmt.Print("请输入您要存款的金额：")
			fmt.Scanln(&amount)
			account.Deposit(password, amount)
		case 2:
			fmt.Print("您正在办理的业务是取款，请输入密码：")
			fmt.Scanln(&password)
			fmt.Print("请输入您要取款的金额：")
			fmt.Scanln(&amount)
			account.WithDrown(password, amount)
		case 3:
			fmt.Print("您正在办理的业务是查询，请输入密码：")
			fmt.Scanln(&password)
			account.Query(password)
		default:
			return
		}

	}

}

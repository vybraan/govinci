package main

import (
	"govinci/core"
	"govinci/htmlout"
	"os"
)

func App(ctx *core.Context) core.View {
	return core.SafeArea(
		core.Scroll(
			core.Column(
				WalletHeader(),
				core.Spacer(16),
				BalanceCard(),
				core.Spacer(20),
				ActionRow(),
				core.Spacer(28),
				SectionTitle("Recent Transactions"),
				TransactionList(),
				core.Spacer(40),
			),
		),
	)
}
func WalletHeader() core.View {
	return core.Row(
		core.Image("https://dummyimage.com/60x60/009688/ffffff&text=G"),
		core.Spacer(12),
		core.Column(
			core.Text("Govinci Wallet", core.FontSize(20), core.FontWeight(core.Bold)),
			core.Text("Welcome back, Ismael", core.TextColor("#555")),
		),
	)
}
func BalanceCard() core.View {
	return core.Card(
		core.Column(
			core.Text("Available Balance", core.TextColor("#8E8E93"), core.FontSize(13)),
			core.Spacer(4),
			core.Text("MZN 42,750.00", core.FontSize(28), core.TextColor("#007AFF"), core.FontWeight(core.Bold)),
		),
	)

}
func MaterialButton(label string, onClick func()) core.View {
	return core.Button(label,
		onClick,
		core.BackgroundColor("#6200EE"),
		core.TextColor("#FFFFFF"),
		core.Padding(12),
		core.BorderRadius(4),
	)
}

func ActionRow() core.View {
	return core.Row(
		MaterialButton("Transfer",
			transferFunc,
		),
		core.Spacer(8),
		MaterialButton("Recharge",
			rechargeFunc,
		))
}
func SectionTitle(label string) core.View {
	return core.Text(label, core.FontSize(18), core.FontWeight(core.Bold), core.TextColor("#222"))
}
func TransactionList() core.View {
	return core.Column(
		TransactionItem("Farmácia", "-750 MZN", "#F44336"),
		TransactionItem("Transferência recebida", "+10,000 MZN", "#4CAF50"),
		TransactionItem("Recarga de saldo", "+3,500 MZN", "#4CAF50"),
	)
}

func TransactionItem(title, amount, color string) core.View {
	return core.Row(
		core.Text(title),
		core.Spacer(12),
		core.Text(amount, core.TextColor(color), core.FontWeight(core.Bold)),
	)
}

func transferFunc() {

}
func rechargeFunc() {

}

func main() {
	ctx := core.NewContext()
	node := App(ctx).Render(ctx)
	output := htmlout.ExportHTML(node)

	// Escrever para o ficheiro
	err := os.WriteFile("hello.html", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}

package example

func ExampleClient_Write() {
	client := newClient()
	client.Write()
	// Output: [DEBUG] client clock at: 3
}

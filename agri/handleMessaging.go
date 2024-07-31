package agri

import (
	"agri-mentor/chatbot"
	"encoding/json"
	"net/http"
)

// var message []string

func HandleSendMessage(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodPost {
	// 		// Parse the form data
	// 		r.ParseForm()
	// 		userInput := r.FormValue("userInput")

	// 		// Process the input (e.g., call your AI model here)
	// 		processedMessage := chatbot.ProcessInput(userInput)

	// 		// Append messages to the global slice
	// 		message = append(message, "You: "+userInput)
	// 		message = append(message, "AI: "+processedMessage)
	// 		// Return the processed message
	// 		// fmt.Fprintf(w, "response:%s", processedMessage) // Send the response back

	// 	}
	// }

	if r.Method == http.MethodPost {
		var msg Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process the message and generate a response
		response := chatbot.ProcessInput(msg.Content)

		// Add both the user's message and the AI's response to the messages slice
		messages = append(messages, Message{Content: "User: " + msg.Content})
		messages = append(messages, Message{Content: "AI: " + response})

		// Return the new messages
		json.NewEncoder(w).Encode(messages)
	} else if r.Method == http.MethodGet {
		// Return all messages
		json.NewEncoder(w).Encode(messages)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

# Sentimental chatbot

This allows to create conversation scripts as trees that will flow according to the evaluated sentiment of the user's responses.

The only implemented script retrieves a product review from a user.

## Design

### Core

The core modules are `conversation` and `nlp`. The `conversation` module is responsible for managing the conversation flow and the `nlp` module is responsible for the sentiment analysis.

`Conversation`s are made out of `Interaction`s organized in a tree structure where each interaction presents a statement/question that may follow a positive, neutral or ambivalent answer.

Sentiment analysis can use either a lexicon or a machine learning model to evaluate if the responses can be considered positive, neutral or negative. The definition of the script allows us to set which kind of analysis we want to use on each statement/question. 

Per the tests realized during building, I learned that closed, yes or no questions tend to work better with lexicon analysis, while open questions tend to work better with machine learning models.

### Scripts

Scripts embed `Conversation` and carry any specific information that needs to be collected from the user. 

### Transport

Transport layers need to implement the `Transport` interface, which is made out of a single method `Input(string)string`.

### Data models

`conversation.Interaction` is the most important model, it composes the conversation trees by defining the possible flows, specifies which kind of sentiment analysis should be used and accept functions to extract information from the messages.

`scripts.ReviewConversation` embeds `conversation.Conversation` and add the fields specific to this kind of dialog by embedding `scripts.Review`.

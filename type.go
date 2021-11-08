package main

type RestResponse struct {
    Ok bool `json:"ok"`
    Result []Update `json:"result"`
}

type Update struct {
    UpdateId int    `json:"update_id"`
    Message Message `json:"message"`
    EditMessage Message `json:"edited_message"`
    CallbackQuery CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
    Id string              `json:"id`
    Data string     `json:"data"`
    Message Message `json:"message"`
    From User       `json:"from"`
}

type BotDelMessage struct {
    ChatId int      `json:"chat_id"`
    MessageId int   `json:"message_id"`
}

type Message struct {
    MessageId int `json:"message_id"`
    From User       `json:"from"`
    Date int        `json:"date"`
    Chat Chat       `json:"chat"`

    ForwardFrom User `json:"forward_from"`
    ForwardFromMessageID int `json:"forward_from_message_id"`
    ForwardFromChat Chat `json:"forward_from_chat"`
    Text string     `json:"text"`
    MediaGroupId string   `json:"media_group_id"`
    
    ReplyToMessage ReplyToMessage `json:"reply_to_message"`
    NewChatMembers []User `json:"new_chat_members"`
    LeftChatMember User `json:"left_chat_member"`

    CaptionEntities []Entities `json:"caption_entities"`
    Entities []Entities `json:"entities"`
    Caption string `json:"caption"`
}

type AnswerCallbackQuery struct{
    CallbackQueryId string `json:"callback_query_id"`
    Text string `json:"text"`
    ShowAlert bool `json:"show_alert"`
}
type Entities struct{
    Type string     `json:"type"`
    Offset int      `json:"offset"`
    Length int      `json:"length"`
    Url string      `json:"url"`
    User User       `json:"user"`
}

type User struct {
    Id int              `json:"id`
    IsBot bool         `json:"is_bot"`
    Username string     `json:"username"`
    FirstName string   `json:"first_name"`
    LastName string    `json:"last_name"`
}

type ReplyToMessage struct {
    MessageId int `json:"message_id"`
    From User       `json:"from"`
    Chat Chat       `json:"chat"`
}


type PostResponse struct {
    Ok bool `json:"ok"`
    Result PostUpdate `json:"result"`
}

type PostUpdate struct {
    PostMessageId int `json:"message_id"`
}

type RestrictChatMember struct{
    Chat_id int `json:"chat_id"`
    User_id int `json:"user_id"`
    Permissions ChatPermissions `json:"permissions"`
    UnitDate int `json:"until_date"`
}

type ChatPermissions struct{
    Can_send_messages bool `json:"can_send_messages"`
    Can_send_media_messages bool `json:"can_send_media_messages"`
    Can_send_polls bool `json:"can_send_polls"`
    Can_send_other_messages bool `json:"can_send_other_messages"`
    Can_add_web_page_previews bool `json:"can_add_web_page_previews"`
    Can_invite_users bool `json:"can_invite_users"`
    Can_pin_messages bool `json:"can_pin_messages"`
}

type Chat struct {
    Id int               `json:"id"`
    Username string     `json:"username"`
    First_name string   `json:"first_name"`
    Last_name string    `json:"last_name"`
}

type ForMessage  struct {
    ChatId int `json:"chat_id"`
    FromChatId int `json:"from_chat_id"`
    MessageId int `json:"message_id"`
}

type CopyMessage struct {
    Chat_id int `json:"chat_id"`
    From_chat_id int `json:"from_chat_id"`
    Message_id int `json:"message_id"`

}


type BotSendMessage struct {
    ChatId int          `json:"chat_id"`
    Text string         `json:"text"`
    ParseMode string    `json:"parse_mode"`
}

type InlineKeyboardRow struct {
    InlineKeyboardRow [][]InlineKeyboard `json:"inline_keyboard"`
}

type InlineKeyboard struct {
    Text string     `json:"text"`
    CallData string  `json:"callback_data"`
}


type BotMessage struct {
    ChatId int                      `json:"chat_id"`
    Text string                     `json:"text"`
    MessageId int                   `json:"message_id"`
    ParseMode string                `json:"parse_mode"`
    Entities Entities               `json:"entities"`
    Reply_Markup InlineKeyboardRow  `json:"reply_markup"`
}

type BanChatMember struct {
    ChatId int                      `json:"chat_id"`
    UserId int                      `json:"user_id"`
    RevokeMessages bool             `json:"revoke_messages"`
}

type UnbanChatMember struct {
    ChatId int                      `json:"chat_id"`
    UserId int                      `json:"user_id"`
    OnlyIfBanned bool               `json:"only_if_banned"`
}

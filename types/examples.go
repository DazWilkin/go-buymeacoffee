package types

// Using the examples on the API Documentation
// https://developers.buymeacoffee.com/#/apireference
// Replaced all next_page_url with null
// Because these are URLs and include scheme://host:path they break the test server
// Which wants URLs of the form e.g. http://localhost:8080/
var (
	ExamplePurchase = []byte(`{
		"purchase_id": 2621,
		"purchased_on": "2020-08-05 09:38:26",
		"purchase_updated_on": "2020-08-05 09:38:26",
		"purchase_is_revoked": 0,
		"purchase_amount": "1.00",
		"purchase_currency": "GBP",
		"purchase_question": "share your email?",
		"payer_email": "quipfosssra@gmail.com",
		"payer_name": "Quip Fora",
		"extra": {
		  "reward_id": 1,
		  "reward_title": "AMA",
		  "reward_description": "Ask me anything",
		  "reward_confirmation_message": "Join your slot here on zoom: link",
		  "reward_question": "share your email?",
		  "reward_used": 3,
		  "reward_created_on": "2020-05-19 10:09:29",
		  "reward_updated_on": "2020-05-19 10:09:29",
		  "reward_deleted_on": null,
		  "reward_is_active": 1,
		  "reward_image": "https://cdn.buymeacoffee.com/uploads/project_updates/2020/05/6cf05f963a4896bab8e377f49da20c8c.jpg",
		  "reward_slots": 10,
		  "reward_coffee_price": "1.00",
		  "reward_order": 0
		}
	  }`)
	ExamplePurchases = []byte(`{
		"current_page": 1,
		"data": [
		 
		  {
			"purchase_id": 30,
			"purchased_on": "2020-05-22 13:17:48",
			"purchase_updated_on": "2020-05-22 13:17:48",
			"purchase_is_revoked": 0,
			"purchase_amount": "8.00",
			"purchase_currency": "USD",
			"purchase_question": "What is your zoom id",
			"payer_email": "f783jksazx@privacy-mail.top",
			"payer_name": "",
			"extra": {
			  "reward_id": 28,
			  "reward_title": "Group Yoga on June 1st",
			  "reward_description": "Join the 60-minute group yoga session where we practice Vinyasa. Beginner-friendly. ",
			  "reward_confirmation_message": "Join using this Zoom link:&nbsp;https://us02web.zoom.us/j/89927509138Here is the link to the answer form:&nbsp;https://forms.gle/mjXCVrczAQV53UTD7Here is the link to the Make Me Laugh round:&nbsp;https://forms.gle/YX8wdLyVskiG74dbAThanks for playing!",
			  "reward_question": "What is your zoom id?",
			  "reward_used": 1,
			  "reward_created_on": "2020-05-22 08:29:54",
			  "reward_updated_on": "2020-05-22 08:29:54",
			  "reward_deleted_on": null,
			  "reward_is_active": 1,
			  "reward_image": "https://cdn.buymeacoffee.com/uploads/project_updates/2020/05/5fdf77e7db86de743998100b29697265.jpg",
			  "reward_slots": null,
			  "reward_coffee_price": "0.00",
			  "reward_order": 0
			}
		  },
		  {
			"purchase_id": 12,
			"purchased_on": "2020-05-20 08:27:34",
			"purchase_updated_on": "2020-05-20 08:27:34",
			"purchase_is_revoked": 0,
			"purchase_amount": "1.00",
			"purchase_currency": "USD",
			"purchase_question": "share your email?",
			"payer_email": "cgtwr@gurumail.xyz",
			"payer_name": "",
			"extra": {
			  "reward_id": 1,
			  "reward_title": "AMA",
			  "reward_description": "Ask me anything",
			  "reward_confirmation_message": "Join your slot here on zoom: link",
			  "reward_question": "share your email?",
			  "reward_used": 3,
			  "reward_created_on": "2020-05-19 10:09:29",
			  "reward_updated_on": "2020-05-19 10:09:29",
			  "reward_deleted_on": null,
			  "reward_is_active": 1,
			  "reward_image": "https://cdn.buymeacoffee.com/uploads/project_updates/2020/05/6cf05f963a4896bab8e377f49da20c8c.jpg",
			  "reward_slots": 10,
			  "reward_coffee_price": "1.00",
			  "reward_order": 0
			}
		  }
		],
		"first_page_url": "https://developers.buymeacoffee.com/api/v1/extras?page=1",
		"from": 1,
		"last_page": 1,
		"last_page_url": "https://developers.buymeacoffee.com/api/v1/extras?page=1",
		"next_page_url": null,
		"path": "https://developers.buymeacoffee.com/api/v1/extras",
		"per_page": 5,
		"prev_page_url": null,
		"to": 4,
		"total": 4
	  }`)
	ExampleSubscription = []byte(`{
		"subscription_id": 7979,
		"subscription_cancelled_on": null,
		"subscription_created_on": "2020-06-03 05:25:03",
		"subscription_updated_on": "2020-06-03 05:25:03",
		"subscription_current_period_start": "2020-06-03 05:25:03",
		"subscription_current_period_end": "2025-06-03 05:25:03",
		"subscription_coffee_price": "0.000",
		"subscription_coffee_num": 1,
		"subscription_is_cancelled": null,
		"subscription_is_cancelled_at_period_end": null,
		"subscription_currency": "USD",
		"subscription_message": null,
		"message_visibility": 1,
		"subscription_duration_type": "lifetime-giveaway",
		"referer": null,
		"country": null,
		"transaction_id": "GIVE_AWAY",
		"payer_email": "rfy8jzy0yo@myinbox.icu",
		"payer_name": "rfy8jzy0yo"
	  }
	  `)
	ExampleSubscriptions = []byte(`{
		"current_page": 1,
		"data": [
			{
				"subscription_id": 10647,
				"subscription_cancelled_on": null,
				"subscription_created_on": "2020-07-29 15:15:22",
				"subscription_updated_on": "2020-07-29 15:15:22",
				"subscription_current_period_start": "2020-07-29 15:15:18",
				"subscription_current_period_end": "2020-08-29 15:15:18",
				"subscription_coffee_price": "1.000",
				"subscription_coffee_num": 1,
				"subscription_is_cancelled": null,
				"subscription_is_cancelled_at_period_end": null,
				"subscription_currency": "USD",
				"subscription_message": null,
				"message_visibility": 1,
				"subscription_duration_type": "month",
				"referer": null,
				"country": null,
				"transaction_id": "sub_HjkUcn*******",
				"payer_email": "ronald.thyu@gmail.org",
				"payer_name": "ronald.thomas"
			}
			
		],
		"first_page_url": "https://developers.buymeacoffee.com/api/v1/subscriptions?page=1",
		"from": 1,
		"last_page": 36,
		"last_page_url": "https://developers.buymeacoffee.com/api/v1/subscriptions?page=36",
		"next_page_url": null,
		"path": "https://developers.buymeacoffee.com/api/v1/subscriptions",
		"per_page": 5,
		"prev_page_url": null,
		"to": 5,
		"total": 179
	}`)
	ExampleSupporter = []byte(`{
		"support_id": 245731,
		"support_note": null,
		"support_coffees": 1,
		"transaction_id": "3JV542690H05****",
		"support_visibility": 1,
		"support_created_on": "2020-08-05 09:38:26",
		"support_updated_on": "2020-08-05 09:38:26",
		"transfer_id": null,
		"supporter_name": null,
		"support_coffee_price": "1.0000",
		"support_email": "*****@gmail.com",
		"is_refunded": null,
		"support_currency": "GBP",
		"support_note_pinned": 0,
		"referer": null,
		"country": "IN",
		"payer_email": "****@gmail.com",
		"payment_platform": "paypal",
		"payer_name": "Quip Fora"
	  }`)
	ExampleSupporters = []byte(`{
		"current_page": 1,
		"data": [
		  {
			"support_id": 245731,
			"support_note": null,
			"support_coffees": 1,
			"transaction_id": "3JV542690H0XXX",
			"support_visibility": 1,
			"support_created_on": "2020-08-05 09:38:26",
			"support_updated_on": "2020-08-05 09:38:26",
			"transfer_id": null,
			"supporter_name": null,
			"support_coffee_price": "1.0000",
			"support_email": "****@gmail.com",
			"is_refunded": null,
			"support_currency": "GBP",
			"support_note_pinned": 0,
			"referer": null,
			"country": "IN",
			"payer_email": "*****@gmail.com",
			"payment_platform": "paypal",
			"payer_name": "** Fora"
		  },
		  {
			"support_id": 245730,
			"support_note": null,
			"support_coffees": 1,
			"transaction_id": "7S721018AR89XXXX",
			"support_visibility": 1,
			"support_created_on": "2020-08-05 09:34:43",
			"support_updated_on": "2020-08-05 09:34:43",
			"transfer_id": null,
			"supporter_name": null,
			"support_coffee_price": "3.0000",
			"support_email": "*****@gmail.com",
			"is_refunded": null,
			"support_currency": "GBP",
			"support_note_pinned": 0,
			"referer": null,
			"country": "IN",
			"payer_email": "*****@gmail.com",
			"payment_platform": "paypal",
			"payer_name": "Quip Fora"
		  },
		  {
			"support_id": 243815,
			"support_note": null,
			"support_coffees": 1,
			"transaction_id": "7A1953730C******",
			"support_visibility": 1,
			"support_created_on": "2020-08-03 10:50:58",
			"support_updated_on": "2020-08-03 10:50:58",
			"transfer_id": null,
			"supporter_name": null,
			"support_coffee_price": "3.0000",
			"support_email": "**@****.com",
			"is_refunded": null,
			"support_currency": "GBP",
			"support_note_pinned": 0,
			"referer": null,
			"country": null,
			"payer_email": "***@***.com",
			"payment_platform": "paypal",
			"payer_name": "asdas"
		  },
		  {
			"support_id": 233527,
			"support_note": null,
			"support_coffees": 1,
			"transaction_id": "60E21623W*****",
			"support_visibility": 1,
			"support_created_on": "2020-07-23 04:12:28",
			"support_updated_on": "2020-07-23 04:12:28",
			"transfer_id": null,
			"supporter_name": null,
			"support_coffee_price": "3.0000",
			"support_email": "****@buymeacoffee.com",
			"is_refunded": null,
			"support_currency": "GBP",
			"support_note_pinned": 0,
			"referer": null,
			"country": null,
			"payer_email": "*****@buymeacoffee.com",
			"payment_platform": "paypal",
			"payer_name": ""
		  },
		  {
			"support_id": 232695,
			"support_note": null,
			"support_coffees": 1,
			"transaction_id": "1GW08219017XXXX",
			"support_visibility": 1,
			"support_created_on": "2020-07-22 04:13:49",
			"support_updated_on": "2020-07-22 04:13:49",
			"transfer_id": null,
			"supporter_name": null,
			"support_coffee_price": "3.0000",
			"support_email": "****@brew.com",
			"is_refunded": null,
			"support_currency": "GBP",
			"support_note_pinned": 0,
			"referer": null,
			"country": null,
			"payer_email": "***@brew.com",
			"payment_platform": "paypal",
			"payer_name": "jijo sunny"
		  }
		],
		"first_page_url": "https://developers.buymeacoffee.com/api/v1/supporters?page=1",
		"from": 1,
		"last_page": 2,
		"last_page_url": "https://developers.buymeacoffee.com/api/v1/supporters?page=2",
		"next_page_url": null,
		"path": "https://developers.buymeacoffee.com/api/v1/supporters",
		"per_page": 5,
		"prev_page_url": null,
		"to": 5,
		"total": 10
	  }`)
)

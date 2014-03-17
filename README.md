# JSON benchmarking

Data type: Twitter streaming API filter / sample endpoint, JSON format.

Some lines look like this:

    {
      "created_at": "Mon Mar 17 04:46:02 +0000 2014",
      "id": 445420773839503360,
      "id_str": "445420773839503360",
      "text": "I'm pissed I have to get my nails taken off, F that chicken bruh.",
      "source": "<a href=\"http://twitter.com/download/iphone\" rel=\"nofollow\">Twitter for iPhone</a>",
      "truncated": false,
      "in_reply_to_status_id": null,
      "in_reply_to_status_id_str": null,
      "in_reply_to_user_id": null,
      "in_reply_to_user_id_str": null,
      "in_reply_to_screen_name": null,
      "user": {
        "id": 236481161,
        "id_str": "236481161",
        "name": "VNasty❤️",
        "screen_name": "Vetraa_",
        "location": "Chasing Dreams..",
        "url": null,
        "description": "God is my number one. I'm loquacious and somewhat sagacious. Not Dr. Phil but I let them know it's real. January 5th ♑️",
        "protected": false,
        "followers_count": 361,
        "friends_count": 199,
        "listed_count": 0,
        "created_at": "Mon Jan 10 17:53:52 +0000 2011",
        "favourites_count": 940,
        "utc_offset": -21600,
        "time_zone": "Mountain Time (US & Canada)",
        "geo_enabled": true,
        "verified": false,
        "statuses_count": 21904,
        "lang": "en",
        "contributors_enabled": false,
        "is_translator": false,
        "is_translation_enabled": false,
        "profile_background_color": "642D8B",
        "profile_background_image_url": "http://abs.twimg.com/images/themes/theme11/bg.gif",
        "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme11/bg.gif",
        "profile_background_tile": true,
        "profile_image_url": "http://pbs.twimg.com/profile_images/445084056326111232/natUutmE_normal.jpeg",
        "profile_image_url_https": "https://pbs.twimg.com/profile_images/445084056326111232/natUutmE_normal.jpeg",
        "profile_link_color": "FF0000",
        "profile_sidebar_border_color": "65B0DA",
        "profile_sidebar_fill_color": "7AC3EE",
        "profile_text_color": "9115B0",
        "profile_use_background_image": true,
        "default_profile": false,
        "default_profile_image": false,
        "following": null,
        "follow_request_sent": null,
        "notifications": null
      },
      "geo": null,
      "coordinates": null,
      "place": null,
      "contributors": null,
      "retweet_count": 0,
      "favorite_count": 0,
      "entities": {
        "hashtags": [],
        "symbols": [],
        "urls": [],
        "user_mentions": []
      },
      "favorited": false,
      "retweeted": false,
      "filter_level": "medium",
      "lang": "en"
    }

Some look like this (retweets):

    {
      "created_at": "Mon Mar 17 04:46:02 +0000 2014",
      "id": 445420773864251400,
      "id_str": "445420773864251392",
      "text": "RT @DionneLister: Awesome Aussie romance author, Ann B Harrison Author has a sexy new book trailer out. Check it out! http://t.co/E6oE62Dbsa",
      "source": "web",
      "truncated": false,
      "in_reply_to_status_id": null,
      "in_reply_to_status_id_str": null,
      "in_reply_to_user_id": null,
      "in_reply_to_user_id_str": null,
      "in_reply_to_screen_name": null,
      "user": {
        "id": 1520143784,
        "id_str": "1520143784",
        "name": "Voice Of Indie",
        "screen_name": "VoiceOfIndie",
        "location": "Lansing, Michigan USA",
        "url": "http://www.goodreads.com/author/show/6112818.Beem_Weeks",
        "description": "Supporting indie authors, musicians, artists, and websites. Author of the novel Jazz Baby. Follow me, or the puppy gets it.",
        "protected": false,
        "followers_count": 6757,
        "friends_count": 7355,
        "listed_count": 185,
        "created_at": "Sat Jun 15 18:46:08 +0000 2013",
        "favourites_count": 6573,
        "utc_offset": -10800,
        "time_zone": "Atlantic Time (Canada)",
        "geo_enabled": false,
        "verified": false,
        "statuses_count": 24961,
        "lang": "en",
        "contributors_enabled": false,
        "is_translator": false,
        "is_translation_enabled": false,
        "profile_background_color": "C0DEED",
        "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
        "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
        "profile_background_tile": false,
        "profile_image_url": "http://pbs.twimg.com/profile_images/344513261583113162/24510cdf3f4ba296e78b8918b9c1dbf3_normal.jpeg",
        "profile_image_url_https": "https://pbs.twimg.com/profile_images/344513261583113162/24510cdf3f4ba296e78b8918b9c1dbf3_normal.jpeg",
        "profile_link_color": "0084B4",
        "profile_sidebar_border_color": "C0DEED",
        "profile_sidebar_fill_color": "DDEEF6",
        "profile_text_color": "333333",
        "profile_use_background_image": true,
        "default_profile": true,
        "default_profile_image": false,
        "following": null,
        "follow_request_sent": null,
        "notifications": null
      },
      "geo": null,
      "coordinates": null,
      "place": null,
      "contributors": null,
      "retweeted_status": {
        "created_at": "Mon Mar 17 04:35:12 +0000 2014",
        "id": 445418047507951600,
        "id_str": "445418047507951616",
        "text": "Awesome Aussie romance author, Ann B Harrison Author has a sexy new book trailer out. Check it out! http://t.co/E6oE62Dbsa",
        "source": "<a href=\"http://www.facebook.com/twitter\" rel=\"nofollow\">Facebook</a>",
        "truncated": false,
        "in_reply_to_status_id": null,
        "in_reply_to_status_id_str": null,
        "in_reply_to_user_id": null,
        "in_reply_to_user_id_str": null,
        "in_reply_to_screen_name": null,
        "user": {
          "id": 382490091,
          "id_str": "382490091",
          "name": "Dionne Lister",
          "screen_name": "DionneLister",
          "location": "Sydney",
          "url": "http://dionnelisterwriter.wordpress.com",
          "description": "Author, editor, copywriter. My fantasy book http://www.amazon.com/Dionne-Lister/e/B007XRT496/ref=ntt_dp_epwb",
          "protected": false,
          "followers_count": 10385,
          "friends_count": 7593,
          "listed_count": 474,
          "created_at": "Fri Sep 30 04:58:07 +0000 2011",
          "favourites_count": 899,
          "utc_offset": -36000,
          "time_zone": "Hawaii",
          "geo_enabled": false,
          "verified": false,
          "statuses_count": 71626,
          "lang": "en",
          "contributors_enabled": false,
          "is_translator": false,
          "is_translation_enabled": false,
          "profile_background_color": "C0DEED",
          "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
          "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
          "profile_background_tile": false,
          "profile_image_url": "http://pbs.twimg.com/profile_images/378800000272566855/f548ce55b74f8e48520e8515181c1221_normal.jpeg",
          "profile_image_url_https": "https://pbs.twimg.com/profile_images/378800000272566855/f548ce55b74f8e48520e8515181c1221_normal.jpeg",
          "profile_link_color": "0084B4",
          "profile_sidebar_border_color": "C0DEED",
          "profile_sidebar_fill_color": "DDEEF6",
          "profile_text_color": "333333",
          "profile_use_background_image": true,
          "default_profile": true,
          "default_profile_image": false,
          "following": null,
          "follow_request_sent": null,
          "notifications": null
        },
        "geo": null,
        "coordinates": null,
        "place": null,
        "contributors": null,
        "retweet_count": 7,
        "favorite_count": 0,
        "entities": {
          "hashtags": [],
          "symbols": [],
          "urls": [
            {
              "url": "http://t.co/E6oE62Dbsa",
              "expanded_url": "http://fb.me/1hamW3aXZ",
              "display_url": "fb.me/1hamW3aXZ",
              "indices": [
                100,
                122
              ]
            }
          ],
          "user_mentions": []
        },
        "favorited": false,
        "retweeted": false,
        "possibly_sensitive": false,
        "lang": "en"
      },
      "retweet_count": 0,
      "favorite_count": 0,
      "entities": {
        "hashtags": [],
        "symbols": [],
        "urls": [
          {
            "url": "http://t.co/E6oE62Dbsa",
            "expanded_url": "http://fb.me/1hamW3aXZ",
            "display_url": "fb.me/1hamW3aXZ",
            "indices": [
              118,
              140
            ]
          }
        ],
        "user_mentions": [
          {
            "screen_name": "DionneLister",
            "name": "Dionne Lister",
            "id": 382490091,
            "id_str": "382490091",
            "indices": [
              3,
              16
            ]
          }
        ]
      },
      "favorited": false,
      "retweeted": false,
      "possibly_sensitive": false,
      "filter_level": "medium",
      "lang": "en"
    }

And the rest look like this (deletes):

    {
      "delete": {
        "status": {
          "id": 445420060824567800,
          "user_id": 112371863,
          "id_str": "445420060824567808",
          "user_id_str": "112371863"
        }
      }
    }


## Objective

Minimize the time it takes to parse 1,000,000 of lines of these types of objects, with the end result being:

* A count of the number of deletes

        cat twitter.json | grep '^{"delete"' | wc -l

* A frequency table of the non-standard background colors, i.e., the `.user.profile_background_color` value where it's different from:

        "profile_background_color": "C0DEED"

* Just to make sure you aren't regexing, the number of entities in the root tweets (not the referenced retweet), summed over all tweets in the set. I.e., the length of each of:
    - `.entities.hashtags`
    - `.entities.symbols`
    - `.entities.urls`
    - `.entities.user_mentions`


## Data

Pull from the sample endpoint until you get 1M lines:

    https://dev.twitter.com/docs/api/1.1/get/statuses/sample

In my dataset, this is about 2.9G uncompressed.

I also ended up with an empty line, somehow, so there's a check that each line is unempty.
Such is data: dirty.

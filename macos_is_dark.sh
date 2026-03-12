#!/bin/bash

function is_dark() {
  osascript -l JavaScript -e \
    "Application('System Events').appearancePreferences.darkMode.get()"
}

# https://github.com/pndurette/zsh-lux/blob/main/zsh-lux.plugin.zsh#L125C1-L137C2
# function example() {
#   local dark_mode=$(dark_flag)
# 
#   if   [[ "$dark_mode" == "true" ]];  then return 0
#   elif [[ "$dark_mode" == "false" ]]; then return 1
#   else
#     return 2
#   fi
# }

echo -n $(is_dark)

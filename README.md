# todo_cli
A cli to perform CRUD operations on your personal todos. 
The state is stored in the root file but that can be changed by changing the path in `storage.go`. The state is stored in json format. 
This is how a todo table looks like :
<img width="735" alt="Screenshot 2024-10-26 at 6 52 41 PM" src="https://github.com/user-attachments/assets/934a9543-3391-4ea6-bf4a-4525418eb302">

## installation 
just clone the repo and cd into it

```zsh
//this will list all the possible commands
  go run ./ 
```

### build into binary
```zsh
  go build ./
```

- then shift the binary to your desired path
- add the path to your rc file to access it from anywhere
- cheerios!

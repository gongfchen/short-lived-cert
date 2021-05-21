package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Use Sha256 to hash an input string
func hash(s string) [32]uint8 {
	data := []byte(s)
	h := sha256.Sum256(data)
	return h
}

func main() {
	fmt.Println("Listening to request...")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			day_num := words[1]
			domain_name := words[0]

			fmt.Println("Hashing certificate and exporting...")
			root := "../../middle-daemon-website-daemon-storage/Daily Certificates/"
			folder := root + domain_name

			cert, err := ioutil.ReadFile(folder + "/certday" + day_num + ".pem")
			check(err)
			cert_text := string(cert)

			hash := hash(cert_text)

			cert_hash := hex.EncodeToString(hash[:])

			f, err := os.Create("../storage/Daily Cert Verification/cert_hash.txt")
			check(err)

			l, err := f.WriteString(cert_hash)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			fmt.Println(l, "bytes written successfully --- Certificate hash written to storage")
			err = f.Close()
			check(err)

			fmt.Println("Verifying proof...")
			ex, err := os.Executable()
			check(err)
			exPath := filepath.Dir(ex)
			file := exPath + "/WDverify.py"
			c := exec.Command("python", file)
			if err := c.Run(); err != nil {
				check(err)
			}
		}

		if scanner.Err() != nil {
			// Handle error.
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

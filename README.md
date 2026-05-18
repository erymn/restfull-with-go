# restfull-with-go - Development Docs

Step yang dilakukan:

1. Initiate project directory
   
   ```bash
   mkdir my-go-project && cd my-go-project
   ```

2. Initialized Go Module
   
   ```bash
   go mod init <mod name>
   go mod init github.com/erymn/my-go-project
   ```
   
   akan tercipta file baru bernama `go.mod` yang berisi mengenai dependensi manifest dan versi module yang digunakan.
   
   Sebagai contoh, jika module  yang diinisialisasi menggunakan format dibawah:
   
   ```bash
   module github.com/erymn/my-go-project
   ```
   
   Deskripsi dari module diatas:
   
   - setiap package yang ada didalam module ini akan direferensikan dengan awalan `github.com/erymn/my-go-project`.
   
   - Package utama di dalam module ini (yang biasanya memakai nama `main.go`) akan dianggap sebagai` root package`.

3. Update Dependency
   
   Untuk update dependency didalam project bisa menggunakan command/perintah
   
   ```bash
   go get <package name>
   ```
   
   contoh penggunaan:
   
   ```bash
   go get example.com/mypackage
   ```
   
   yang terjadi, perintah `go get` akan melakukan update pada file `go.mod` dan `go.sum` dengan informasi `package dependency` terupdate.
   
   - `go.mod` : sebuah file yang berada di root project yang akan merekam semua file dependency
   
   - `go.sum` : bersama dengan file `go.mod`, yang memastikan integritas semua modules dengan menyimpan informasi checksums.

4. Tidy and Vendor Prune
   
   `go mod tidy` adalah perintah yang akan menghapus dependency dari file `go.mod` yang tidak lagi digunakan didalam suatu project.
   
   `go mod vendor` akan populate direktory vendor dengan dependency secara eksplisit di dalam kodenya.

5. Memahami Module Go
   
   1. Menambah 3rd Party Package
      
      ```bash
      go get github.com/package-author/package-name
      ```
      
      Impor di code
      
      ```bash
      import "github.com/package-author/package-name"
      ```
   
   2. Manage Versi Package
      
      Secara default, `go get` akan mengambil yang paling latest.
      
      Versi Spesifik: menggunakan `@` keyword.
      
      ```bash
      go get github.com/package-author/package-name@v1.2.3
      ```
   
   3. Update Packages
      
      Update:
      
      ```bash
      go get -u github.com/package-author/package-name              
      ```
      
      Remove:
      
      ```bash
      go mod tidy          
      ```

## Go Syntax & Concepts

### Variables

- Basic declaration
  
  ```go
  var a int       // declares an integer variable 'a'
  var b string    // declares a string variable 'b'
  ```

- Short declaration
  
  ```go
  a := 10         // declares an integer variable 'a' and initializes it to 10
  b := "hello"    // declares a string variable 'b' and initializes it to "hello"
  ```

- Multiple variables
  
  ```go
  var a, b, c int             // multiple variables of the same type
  var d, e = 400, f = "abc"   // multiple variables with initialization
  ```

### Constants

- Basic declaration
  
  ```go
  const Pi = 3.14
  ```

- Typed constants
  
  ```go
  const Pi float64 = 3.14
  ```

- Multiple constants
  
  ```go
  const ( StatusOK = 200 NotFound = 404 )
  ```

### Data Type in Go

- Basic Type
  
  - Integers
    
    - Signed Integers: `int, int8, int16, int32, int64`
      
      - int: yang paling umum digunakan, ukuran tergantung platform (biasanya 32 dan 64 bits)
      
      - int8, int16, int32, int64, mewakili nilai 8, 16, 32 dan 64 bit integer
    
    - Unsigned Integers: `uint, uint8, uint16, uint32, uint64`
      
      - Sama dengan signed integers, bedanya unsigned tidak bisa menerima data negatif
      
      - `-uint8` juga dikenal sebagai byte
    
    - Machine Dependent Types
      
      - adalah bilangan bulat unsigned yang cukup besar untuk menyimpan bit yang tidak diinterpretasikan.
  
  - Floating Numbers: `float32, float64`
    
    - Mempresentasilan bilangan floating-point presisi tunggal (32 bit) dan ganda (64 bit) 
  
  - Complex Numbers: `complex64, complex128`
    
    - Digunakan untuk bilangan kompleks riil dan imajiner `float32` dan `float64`.
  
  - String
    
    - Merepresentasikan urutan byte (biasanya text). Di `Go`, string immutable/tidak bisa diubah 
  
  - Boolean

- Aggregate Type
  
  Aggregate di Go adalah tipe data gabungan yang menggabungkan beberapa nilai menjadi 1 entitas tunggal.
  
  - Array Type, adalah urutan elemen berukuran tetap dari satu tipe data (data type). Panjang array merupakan bagian dari tipenya
    
    ```go
    var i = [14]int
    ```
  
  - Struct Type, sebuah struct akan mengelompokkan variable (nantinya akan disebut field) dibawah 1 nama, dengan setiap field memiliki nama dan tipe data sendiri. ini sangat berguna untuk mengelompokkan data yang terkait.
    
    ```go
    type Vertex struct {
        X int
        Y int
        Z string
    } 
    ```

- Reference Type
  
  Tipe ini ada tipe yang merujuk ke struktur data yang sudah ada atau mendasarinya dan bukan menyimpan data secara langsung. Biasanya digunakan untuk proses manipulasi dan berbagi data di berbagai bagian program.
  
  - Pointer Type
    
    Pointer menyimpan alamat memori variabel. Pointer memungkinkan manipulasi tidak langsung terhadap nilai pada alamat yang dirujuk.
    
    ```go
    var p *int
    ```
  
  - Slice Type
    
    Slice mirip dengan array tetapi ukurannya dinamis dan fleksibel, sedangkan array ukurannya fixed/pas/tidak dapat diubah. Sebuah slice menunjuk ke sebuah array, memberikan tampilan elemen-elemennya.
    
    ```go
    var i = []int
    ```
  
  - Map Type
    
    Map adalah penyimpanan key-value di mana kuncinya unik. Map digunakan untuk mencari nilai berdasarkan kuncinya.
    
    ```go
    var m map[string]int
    ```
  
  - Channel Type
    
    Channel menyediakan cara bagi goroutine untuk berkomunikasi dan melakukan sinkronisasi. Channel digunakan secara luas dalam pemrograman Go konkuren.
    
    ```go
    var intChan = chan int
    ```

- Interface Type
  
  Sebuah Interface mendefinisikan sekumpulan metode. Setiap tipe/class yang mengimplementasikan metode-metode ini harus memenuhi jumlah method yang ada di Interface. interface ini adalah inti dari sistem dan polimorfisme di dalam Go.
  
  ```go
  type InterfaceName interface {
      Method1() returnType
      Method2(paramType) returnType
  } 
  ```

- Function Type
  
  Adalah first-class value di Go. Sebuah tipe function didefinisikan oleh signaturenya termasuk didalamnya parameter dan nilai kembaliannya.
  
  ```go
  func functionName(parameters) returnType {
  // Function body
  }
  ```

## Mathematical & Logical Operator

- Arithmetic Operator
  
  - Penambahan (`+`)
  
  - Pengurangan (`-`)
  
  - Perkalian (`*`)
  
  - Pembagian (`/`)
  
  - Modulus (`%`), mengembalikan sisa nilai dari hasil pembagian 2 operan

- Assignment Operator
  
  - Simple Assignment (`=`), menetapkan nilai operan disebelah kanan untuk operan disebelah kiri
  
  - Compound Assignment, gabungan arithmetic dan assignment operator seperti (`+=, -=,*=, /=, %=`)

- Comparison Operator
  
  - Equal (`==`): Checks if two values are equal.
  
  - Not Equal (`!=`) : Checks if two values are not equal.
  
  - Greater Than (`>`) : Checks if the left value is greater than the right.
  
  - Less Than (`<`) : Checks if the left value is less than the right.
  
  - Greater Than or Equal To (`>=`) : Checks if the left value is greater than or equal
    to the right.
  
  - Less Than or Equal To (`<=`) : Checks if the left value is less than or equal to the
    right.

- Logical Operator
  
  - Logical AND (`&&`): Returns true if both operands are true.
  
  - Logical OR (`||`): Returns true if either of the operands is true.
  
  - Logical NOT (`!`): Negates the truth value of the operand.

- Bitwise Operator
  
  - AND (`&`): Performs a bitwise AND on two operands.
  
  - OR (`|`) : Performs a bitwise OR on two operands.
  
  - XOR (`^`) : Performs a bitwise XOR on two operands.
  
  - NOT (`~`) : Negates each bit of the operand.
  
  - Left Shift (`<<`) : Shifts the bits of the left operand to the left.
  
  - Right Shift (`>>`) : Shifts the bits of the left operand to the right.

- Increment & Decrement Operator
  
  - Increment (`++`) : Increases the integer value of the operand by one.
  
  - Decrement (`--`) : Decreases the integer value of the operand by one.

- Miscellaneous Operator
  
  - Address Operator (`&`) : Returns the memory address of its operand.
  
  - Pointer Indirection Operator (`*`) : Accesses the value in the location pointed to
    by the operand.

## Control Structure

- `if` and `if-else` statement
  
  `if` digunakan untuk eksekusi blok kode jika suatu kondisi = true/benar
  
  ```go
  if condition {
  // code to execute if the condition is true
  }
  ```
  
  Statement `if` juga bisa memiliki blok else jika kondisi pertama yang ditemukan bernilai = false/salah
  
  ```go
  if condition {
  // code if the condition is true
  } else {
  // code if the condition is false
  }
  ```
  
  Selain itu, pernyataan `if` dapat dimulai dengan pernyataan singkat yang akan dieksekusi sebelum kondisi. Variabel yang dideklarasikan dalam pernyataan ini hanya berada dalam cakupan hingga akhir pernyataan `if`.
  
  ```go
  if a := computeValue(); a > threshold {
  // code to execute if the condition is true
  }
  ```

- Switch statement
  
  Statement ini digunakan untuk menyederhanakan kondisional yang komplek.
  
  ```go
  switch expression {
      case value1: // code to execute if expression == value1
      case value2: // code to execute if expression == value2
      default: // code to execute if no case matches
  }
  ```

- Loops
  
  - for loop, go hanya memiliki 1 looping construct.
    
    ```go
    for condition {
    // code to execute while the condition is true
    }
    ```
  
  - classic for loop, sama seperti C dan Java, dengan initializer, condition dan post statement
    
    ```go
    for init; condition; post {
    // code to execute
    }
    ```
  
  - for-range loop, digunakan untuk iterasi dari suatu element di `slice, map, string `atau `channel`
    
    ```go
    for key, value := range collection {
    // code to execute
    }
    ```
  
  - infinite loop, loop tiada akhir
    
    ```go
    for {
    // code to execute forever
    }
    ```
  
  - break, digunakan di Go untuk keluar dari looping
    
    ```go
    for i := 0; i < 10; i++ {
        if i == 5 {
          break // Exits the loop when i equals 5
        }
        fmt.Println(i)
    }
    ```
  
  - continue, digunakan untuk skip disuatu porsi dari suatu loop dan akan memulai dengan loop baru
    
    ```go
    for i := 0; i < 10; i++ {
        if i % 2 == 0 {
          continue // Skips the rest of the loop body for even numbers
        }
        fmt.Println(i) // Only odd numbers are printed
    }
    ```

## Function in Go

- Define function in Go
  
  didefinisikan dengan format `func functionName (parameter 1 type1, parameter2 type2) returnType {}`.
  
  ```go
  func functionName(parameter1 type1, parameter2 type2) returnType {
      // function body
  }
  ```
  
  ```go
  func functionName() returnType {
      // function body
  }
  ```

- Variadic function
  
  Fungsi dapat bersifat variadik, artinya fungsi tersebut dapat menerima sejumlah argumen variabel dengan tipe tertentu. Fungsi ini sangat berguna ketika Anda membutuhkan fungsi untuk menangani sejumlah input yang tidak ditentukan. Sintaks untuk fungsi variadik melibatkan penambahan tipe sebelum tipe parameter terakhir dalam tanda tangan fungsi.
  
  ```go
  // sum takes any number of integer arguments and returns their sum.
  func sum(nums …int) int {
      total := 0
      for _, num := range nums {
          total += num
      }
      return total
  }
  ```
  
  Cara memanggil fungsi sum diatas
  
  ```go
  func main() {
      fmt.Println(sum(1, 2))          // Output: 3
      fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
  
      // Passing a slice of integers with '…'
      numbers := []int{1, 2, 3, 4, 5, 6}
      fmt.Println(sum(numbers…))    // Output: 21
  }
  ```

- Usage in API Development
  
  Di pengembangan API, fungsi selalu digunakan secara intensif untuk handle HTTP request, peng-kapsulan logika bisnis dan interaksi dengan database. 
  
  Contoh dasar HTTP Server dengan Handler:
  
  ```go
  func homeHandler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Welcome to the Home Page!")
  }
  ```
  
  ```go
  func aboutHandler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "About Us")
  }
  ```
  
  ```go
  func main() {
      http.HandleFunc("/", homeHandler)
      http.HandleFunc("/about", aboutHandler)
      http.ListenAndServe(":8080", nil)
  }
  ```
  
  Penjelasan:
  
  Contoh diatas `homeHandler` dan `aboutHandler `ada fungsi yang menangani HTTP Request di root path url ("`/`") dan path `/about`.



## Penanganan Error dan Strateginya

- Tipe error di Go
  
  - Tipe Error
    
    ```go
    type error interface {
        Error() string
    }
    ```
    
    Go memiliki tipe interface bernama `Error()`. Error ini dapat menyimpan semua nilai error dan `null `error.
  
  - Buat Error
    
    Error standar Go dapat dibuat dengan menggunakan `errors`, fungsi error atau menggunakan class `fmt.Errorf` untuk pesan error yang terformat.
    
    ```go
    err1 := errors.New("error occurred")
    err2 := fmt.Errorf("error occurred at %v", time.Now())
    ```

- Strategi penanganan Error
  
  - Returning Errors/Pengembalian Error
    
    Fungsi yang dapat mengalami kesalahan biasanya mengembalikan kesalahan sebagai nilai kembalian terakhirnya.
    
    Saat memanggil fungsi yang mengembalikan kesalahan, Anda harus memeriksa apakah kesalahan tersebut bukan nil dan menanganinya dengan tepat.
    
    ```go
    func doSomething() (result string, err error) {
        // … function logic …
        if somethingWentWrong {
            return "", errors.New("something went wrong")
        }
        return "success", nil
    }
    ```
  
  - Tipe Error Custom
    
    Kamu juga dapat menggunakan Error Custom dengan meng-implementasikan interface error. Tipe error custom sangat berguna ketika dibutuhkan konteks atau data mengenai error diperlukan.
    
    ```go
    type MyError struct {
        Msg string
        Code int
    }
    ```
    
    ```go
    func (e *MyError) Error() string {
        return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
    }
    ```
  
  - Menangani Error 
    
    Penanganan error yang tepat melibatkan pemeriksaan nilai kembalian kesalahan dan mengambil tindakan yang sesuai.
    
    Praktik umum meliputi pencatatan kesalahan, mengembalikan kesalahan ke atas tumpukan panggilan,
    atau menangani kesalahan dengan menerapkan mekanisme pemulihan.
    
    ```go
    result, err := doSomething()
    if err != nil {
        // handle error, for example, log it and/or return it
        log.Println(err)
        return err
    }
    // proceed with 'result' if no error
    ```
  
  - Panic & Recover
    
    - Go memiliki fungsi panic dan recover bawaan, tetapi penggunaannya berbeda dari pengecualian di bahasa lain.
    
    - Panic menghentikan eksekusi normal fungsi saat ini dan memulai
      panicking, yang melepaskan tumpukan hingga recover dipanggil atau program crash.
    
    - Recover digunakan untuk mendapatkan kembali kendali atas goroutine yang mengalami panic dan melanjutkan eksekusi normal.
    
    - Penggunaan panic dan recover umumnya tidak disarankan kecuali untuk menangani situasi yang benar-benar luar biasa atau kasus di mana eksekusi tidak dapat dilanjutkan.
  
  - Wrapping Error
    
    Mulai Go versi 1.13 wrapping error diperkenalkan, dan memungkinkan untuk menambahkan konteks tambahan dalam suatu error sambil mempertahankan error aslinya/original. Hal ini dilakukan dengan menggunakan kata kerja `%w` dengan `fmt.Errorf`.
    
    ```go
    if err != nil {
        // Wrap the error with additional context
        return fmt.Errorf("failed to do something: %w", err)
    }
    ```

## Structuring Go Code & Go Idiomatic Style

Contoh struktur direktori project go sederhana

```go
/my-go-project
  /bin
  /pkg
  /src
    /github.com
      /username
        /my-go-project
         /cmd
           /myapp
             main.go
           /pkg
             /mypackage
               mypackage.go
           go.mod
           go.sum
```

Package di Go

- Buat package
  
  Package di Go adalah sebuah direktory yang didalamnya terdapat 1 atau lebih file `.go`.
  
  Setiap direktory didalam `src` idealnya mewakili 1 package.

- Penamaan Package
  
  Nama package adalah nama direktori yang memuatnya.

- Exported & UnExported Package
  
  Dalam Go, nama variable, tipe, fungsi, dsb yang diawali dengan huruf kapital/besar akan terlihat dari luar package (public identifier).
  
  Pengidentifikasian nama variable, tipe, fungsi, dsb yang diawali huruf kecil, hanya bisa diakses didalam package itu sendiri.

- Package pihak ketiga
  
  Go memiliki ekosistem paket pihak ketiga yang luas. Anda dapat memanfaatkan paket-paket ini untuk menambahkan fungsionalitas ke aplikasi Anda tanpa harus menciptakan kembali hal yang sudah ada.
  Berhati-hatilah dalam menambahkan dependensi eksternal untuk memastikan proyek Anda tetap ringan dan mudah dikelola.



## Best Practice

Seumpama kalian membuat web apps. Mungkin bisa mengikuti panduan struktur package berikut ini:

- `cmd/myapp` : Contains the main application entry point.

- `pkg/database` : Encapsulates database-related operations.

- `pkg/api`: Contains handlers and logic for the REST API.

- `pkg/auth `: Manages authentication and authorization.
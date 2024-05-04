# CIS202 Final Project: File Management System

**Description:**
In this project, you will build a file management system using Python. The system should be able to traverse a given list of directories and individual files, populate file name, timestamp, and file size information, and store them in an appropriate data structure. Additionally, the system should provide various functionalities to interact with the stored files, such as printing a list of files with full paths, timestamps, and file sizes, moving or deleting files based on user input, detecting and removing duplicate files, and more.

**Requirements:**

1. **File Traversal and Data Population**
   - Implement a function that takes a list of directories and individual files as input.
   - Traverse the given list and populate file name, timestamp, and file size information for each file.
   - Store the collected information in an appropriate data structure (e.g., a list of dictionaries or a custom class). You need to research how to use Python to traverse file directories (using Language Models is allowed).

2. **Printing File Information**
   - Implement a function that prints the list of files as an Excel file (xls or xlsx) with their full paths, timestamps, and file sizes in bytes. One file per row. You need to research how to create an Excel file with Python (using Language Models is allowed).

3. **File Operations (Move and Delete)**
   - Implement a function that takes the previously output Excel file modified with the user's desired action as input. You need to research how to use Python to read an Excel file.
   - The user's desired action should be in the new extra column `action` at the end, where `action` can be `delete` or `moveto:/new/path/to/move/to/`. You need to research how to upload file in Colab notebook, and how to use Python to manipulate files (delete/move).
   - The function should perform the specified actions (move or delete files) based on the user input.
   - The user may not mark a desired action for every row. No mark means to leave the file as is.

4. **Duplicate File Detection**
   - Implement a function that detects potential duplicate files based on file size and timestamp.
   - If potential duplicates are found, the function should further verify duplication by checking the file hash. You need to research how to use Python to compute file hashes (it's okay to call system commands to get hash values and read the system output back into a Python string -- you may research how to do this).
   - The function should create a list of actual duplicated files with their full paths, timestamps, file sizes, and the list of files they duplicate, like `[/path1/file1, /path2/file2]`.

5. **Duplicate File Removal**
   - Implement a function that allows the user to mark files for deletion from the previously printed list of duplicate files.
   - The user will modify the Excel file output from the previous step by adding the new extra column `action` at the end, where `action` value can be `delete` or empty (no action).
   - The function should remove the marked duplicate files from the file system.

6. **Error Handling and Validation**
   - Implement appropriate error handling and validation mechanisms for input data and file operations.
   - Handle scenarios such as invalid file paths, missing files.

7. **Documentation**
   - Provide clear documentation for your code, including function descriptions, parameter explanations, and any assumptions or limitations.
   - Include instructions on how to run and test your implementation.

8. **Testing**
   - Create a comprehensive set of test cases to validate the correctness of your implementation (i.e., you need to write Python code to create directories and random files with different sizes/timestamps and make duplicates -- research how to do this).
   - Test cases should cover various scenarios, including different file structures, file operations, and edge cases.

9. **Efficiency and Performance**
   - Consider the efficiency and performance of your implementation, especially for large file systems or large numbers of files.
   - Analyze the time and space complexities of your algorithms and data structures, and optimize them if necessary.

**Submission:**
Submit the URL of your project notebook. Make sure to include clear instructions on how to run and test your implementation.

**Grading Criteria:**
Your project will be evaluated based on the following criteria:

- Correctness: Your implementation should correctly perform the required functionalities and handle edge cases.
- Efficiency: Your implementation should be efficient in terms of time and space complexity.
- Code Quality: Your code should be well-structured, readable, and follow best practices.
- Documentation: Your documentation should be clear and complete, explaining the purpose, functionality, and usage of your implementation.
- Testing: Your test cases should be comprehensive and cover various scenarios.

Feel free to reach out to me at ~~redacted~~ if you have any questions or need further clarification.

# Response to Section 9

The `Ls` function has to be of $O(n)$, where $n$ is the number of files. It is necessary to loop through every file, and optimization is unlikely for this operation.

The `Excel` function is also $O(kn)$, I included $k$ because it is due to the fixed operations it needs to do to insert the value into the cell for each file. I thought about improving this constant, e.g. making 1 call instead of 3 per file when I am populating the columns, but the module I implemented for the sake of ease does not seem to support this.

The `ExcelMvDel` function I implemented is $O(n)$, since it will loop through every row / every file in the spreadsheet to detect if there is an action that needs to be undertaken. One way to improve this would be to sort the spreadsheet, thus costing in general $O(log(n))$, and then loop through the files that are actually marked with an action, skipping the empty ones. But due to the number of files a user usually deal with on their system, the efficiency provided by golang should be enough.

The `LsDupes` function has a lot of space to improve. Currently it takes $O(n^2)$ time to compare every entry, but it is possible to reduce the comparisions needed by implementing a smarter dedupe algorithm. That will take more time to learn, implement and validate then the current version, but will be helpful for a system with more files.

A good decision I think I made was to only hash the file when the cheaper timestamp and file size comparisons have been made. It should be rather rare for a system to have many files that have the same file size and timestamp at the same time, thus reducing the hashing is desriable. Though I did memoize the hashing for `vf[i]`, because it might be used later.

As for space complexity, it is possible for one to run out memory when creating the verbose file list, but considering the modern system's capacity, I don't think it is necessary to optimize that. When developing an app, one must make a balance between efficiency and development time, especially for a student project like this one that is not focused on super optimization, but rather implmentation of learned methods in programming.

Thank you for reading.
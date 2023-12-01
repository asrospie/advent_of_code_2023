use std::fs;
pub fn read_file_lines(filename: String) -> Result<Vec<String>, std::io::Error> {
    match fs::read_to_string(filename) {
        Ok(contents) => {
            let mut temp_storage: Vec<String> = Vec::new();
            for line in contents.lines() {
                temp_storage.push(line.to_string());
            }
            Ok(temp_storage)
        },
        Err(e) => return Err(e),
    }
}

pub fn read_file_string(filename: String) -> Result<String, std::io::Error> {
    match fs::read_to_string(filename) {
        Ok(contents) => Ok(contents),
        Err(e) => return Err(e),
    }
}

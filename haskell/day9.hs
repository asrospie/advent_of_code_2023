exampleFile = "inputs/example.txt"
inputFile = "inputs/input.txt"

main :: IO ()
main = do
  input <- readFile exampleFile
  putStrLn $ show $ sum $ fmap (predictFuture) (strsToInts $ lines input)
  input <- readFile inputFile
  putStrLn $ show $ sum $ fmap (predictFuture) (strsToInts $ lines input)
  input <- readFile exampleFile
  putStrLn $ show $ sum $ fmap (predictPast) (strsToInts $ lines input)
  input <- readFile inputFile
  putStrLn $ show $ sum $ fmap (predictPast) (strsToInts $ lines input)

strsToInts :: [String] -> [[Int]]
strsToInts input = fmap (strToInts) input

strToInts :: String -> [Int]
strToInts input = fmap (strToInt) (words input)

strToInt :: String -> Int
strToInt s = read s

predictFuture :: [Int] -> Int
predictFuture xs = predictFutureHelper [0] xs 

predictPast :: [Int] -> Int
predictPast xs = predictFutureHelper [0] (reverse xs)


predictFutureHelper :: [Int] -> [Int] -> Int
predictFutureHelper above below =
  let diffList = differences below
  in if allZeros below
    then (last above) + (last below)
    else (last above) + predictFutureHelper below diffList


differences :: [Int] -> [Int]
differences [] = []
differences [_] = []
differences (x:y:xs) = (y - x) : differences (y:xs)

allZeros :: [Int] -> Bool
allZeros [] = True
allZeros (x:xs)
  | x /= 0  = False
  | otherwise = allZeros xs

example filename = do
  input <- readFile filename
  putStrLn $ show $ sum $ fmap (predictFuture) (strsToInts $ lines input)

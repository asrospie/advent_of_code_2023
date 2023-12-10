-- imports
import Data.Char (isDigit)
import Data.Char (isSpace)

exampleFile = "inputs/example_6.txt"
inputFile = "inputs/input_6.txt"

main :: IO ()
main = do
  -- part 1
  input <- readFile exampleFile
  let pre_input = fmap (getNumsFromString) (lines input)
  let zipped = zipWith (\x y -> (x, y)) (head pre_input) (pre_input !! 1)
  print zipped
  let all_winners = fmap (winners) zipped
  print all_winners
  let result = mulList all_winners
  print result
  input <- readFile inputFile
  let pre_input = fmap (getNumsFromString) (lines input)
  let zipped = zipWith (\x y -> (x, y)) (head pre_input) (pre_input !! 1)
  print zipped
  let all_winners = fmap (winners) zipped
  print all_winners
  let result = mulList all_winners
  print result

  -- part 2
  input <- readFile exampleFile
  let nums_strs = fmap (getStrNumsFromString) (lines input)
  let num_strs = fmap (joinStrings) nums_strs
  let nums = fmap (strToInt) num_strs
  let num_tuple = ((head nums), (nums !! 1))
  print $ newWinners num_tuple
  input <- readFile inputFile
  let nums_strs = fmap (getStrNumsFromString) (lines input)
  let num_strs = fmap (joinStrings) nums_strs
  let nums = fmap (strToInt) num_strs
  let num_tuple = ((head nums), (nums !! 1))
  print $ newWinners num_tuple

mulList :: [Int] -> Int
mulList xs = foldl (\acc x -> acc * x) 1 xs

winners :: (Int, Int) -> Int
winners (time, dist) = winnersHelper (time, dist) 1

winnersHelper :: (Int, Int) -> Int -> Int
winnersHelper (time, dist) acc
  | acc == dist = 0
  | (time - acc) * acc > dist = 1 + winnersHelper (time, dist) (acc + 1) 
  | otherwise = winnersHelper (time, dist) (acc + 1)


newWinners :: (Int, Int) -> Int
newWinners (time, dist) = do
  let end = newWinnersTopDown (time, dist) (time - 1)
  let start = newWinnersBottomUp (time, dist) 1
  end - start + 1

newWinnersBottomUp :: (Int, Int) -> Int -> Int
newWinnersBottomUp (time, dist) acc
  | acc == dist = 0
  | (time - acc) * acc > dist = acc
  | otherwise = newWinnersBottomUp (time, dist) (acc + 1)

newWinnersTopDown :: (Int, Int) -> Int -> Int
newWinnersTopDown (time, dist) acc
  | acc == 0 = 0
  | (time - acc) * acc > dist = acc
  | otherwise = newWinnersTopDown (time, dist) (acc - 1)

strsToInts :: [String] -> [[Int]]
strsToInts input = fmap (strToInts) input

strToInts :: String -> [Int]
strToInts input = fmap (strToInt) (words input)

strToInt :: String -> Int
strToInt s = read s

getNumsFromString :: String -> [Int]
getNumsFromString s = fmap (strToInt) (getStrNumsFromString s)

getStrNumsFromString :: String -> [String]
getStrNumsFromString = go ""
  where
    go :: String -> String -> [String]
    go acc [] = if null acc then [] else [acc]
    go acc (x:xs)
      | isDigit x = go (acc ++ [x]) xs
      | null acc = go acc xs
      | otherwise = acc : go "" xs

removeSpaces :: String -> String
removeSpaces = filter (not . isSpace)

joinStrings :: [String] -> String
joinStrings xs = joinStringsHelper xs ""

joinStringsHelper :: [String] -> String -> String
joinStringsHelper [] y = y
joinStringsHelper [x] y = x ++ y 
joinStringsHelper (x:xs) y = x ++ joinStringsHelper xs y



(ns solution
  (:require [clojure.string :as str]))

(defn floor-div [a b]
  (int (Math/floor (/ (double a) (double b)))))

(defn solve [input]
  (let [lines (str/split-lines input)]
    (loop [rotations lines
           position 50
           total-zero-count 0]
      (if (empty? rotations)
        total-zero-count
        (let [line (first rotations)
              direction (first line)
              distance (Integer/parseInt (subs line 1))
              zeros-this-turn (if (= direction \R)
                                (- (floor-div (+ position distance) 100) (floor-div position 100))
                                (- (floor-div (- position 1) 100) (floor-div (- position distance 1) 100)))
              new-position (if (= direction \R)
                             (+ position distance)
                             (- position distance))]
          (recur (rest rotations)
                 new-position
                 (+ total-zero-count zeros-this-turn)))))))

(defn -main [& args]
  (let [input (slurp "day1/part1/input1.txt")]
    (println (str "The password is: " (solve input)))))

(-main)


(ns solution
  (:require [clojure.string :as str]))

(defn solve [input]
  (let [lines (str/split-lines input)]
    (loop [lines lines
           position 50
           zero-count 0]
      (if (empty? lines)
        zero-count
        (let [line (first lines)
              direction (first line)
              distance (Integer/parseInt (subs line 1))
              new-position (if (= direction \R)
                             (+ position distance)
                             (- position distance))
              wrapped-position (mod new-position 100)]
          (recur (rest lines)
                 wrapped-position
                 (if (zero? wrapped-position)
                   (inc zero-count)
                   zero-count)))))))

(defn -main [& args]
  (let [input (slurp "day1/part1/input1.txt")]
    (println (str "The password is: " (solve input)))))

;; This allows the script to be executed directly
(-main)

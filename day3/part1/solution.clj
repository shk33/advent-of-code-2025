(ns solution
  (:require [clojure.string :as str]))

(defn- get-max-joltage-for-line [line]
  (let [digits (vec line)]
    (->> (for [i (range (count digits))
               j (range (inc i) (count digits))]
           (str (get digits i) (get digits j)))
         (map #(Integer/parseInt %))
         (apply max 0))))

(defn -main [& args]
  (let [input-path (str (.getParent (java.io.File. *file*)) "/input.txt")
        lines (str/split-lines (slurp input-path))
        total-joltage (->> lines
                           (map str/trim)
                           (filter not-empty)
                           (map get-max-joltage-for-line)
                           (reduce +))]
    (println (str "The total output joltage is: " total-joltage))))

(-main)

(ns solution
  (:require [clojure.string :as str]))

(defn find-max-subsequence [line k]
  (if (< (count line) k)
    (str/join (repeat k "0"))
    (loop [i 0
           current-start-index 0
           result-strings []]
      (if (= i k)
        (str/join result-strings)
        (let [remaining-to-find (- k i)
              search-end-index (- (count line) remaining-to-find)
              search-str (if (<= current-start-index search-end-index)
                           (subs line current-start-index (inc search-end-index))
                           "")]
          (if (empty? search-str)
            (str/join (concat result-strings (repeat remaining-to-find "0")))
            (let [best-digit-str (last (sort (map str search-str)))
                  best-digit-offset (.indexOf search-str best-digit-str)
                  best-digit-abs-index (+ current-start-index best-digit-offset)]
              (recur (inc i)
                     (inc best-digit-abs-index)
                     (conj result-strings best-digit-str)))))))))

(defn -main [& args]
  (let [input-path (str (.getParent (java.io.File. *file*)) "/input.txt")
        k 12
        lines (str/split-lines (slurp input-path))
        total-joltage (->> lines
                           (map str/trim)
                           (filter not-empty)
                           (map #(find-max-subsequence % k))
                           (map bigint)
                           (reduce +))]
    (println (str "The total output joltage is: " total-joltage))))

(-main)

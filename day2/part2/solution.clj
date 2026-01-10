
(ns solution
  (:require [clojure.string :as str]
            [clojure.set :as set]))

(defn parse-ranges [input-str]
  (let [range-strs (str/split (str/trim input-str) #",")]
    (reduce (fn [[ranges max-id] r-str]
              (let [parts (str/split r-str #"-")
                    start (bigint (first parts))
                    end (bigint (second parts))]
                [(conj ranges [start end]) (max max-id end)]))
            [[] 0]
            range-strs)))

(defn generate-invalid-ids [max-id]
  (loop [base-num 1
         invalid-ids []]
    (let [base-s (str base-num)
          first-repetition-val (bigint (str base-s base-s))]
      (if (> first-repetition-val max-id)
        invalid-ids ; Optimization: smallest repetition is too big, stop outer loop
        (let [generated-for-base (loop [current-repeated-s base-s
                                         acc []]
                                   (let [n-invalid (bigint (str current-repeated-s base-s))] ; Creates next repetition
                                     (if (> n-invalid max-id)
                                       acc ; Break inner loop, return accumulated
                                       (recur (str current-repeated-s base-s) (conj acc n-invalid)))))
              all-new-invalid-ids (into invalid-ids generated-for-base)] ; Add all generated for this base-s
          (recur (inc base-num) all-new-invalid-ids))))))

(defn solve [input-str]
  (let [[ranges max-id] (parse-ranges input-str)
        potential-ids (generate-invalid-ids max-id)
        found-ids (loop [ids potential-ids
                         found (transient #{})]
                    (if-let [id (first ids)]
                      (if (some #(<= (first %) id (second %)) ranges)
                        (recur (rest ids) (conj! found id))
                        (recur (rest ids) found))
                      (persistent! found)))]
    (reduce + found-ids)))

(defn main []
  (let [input-str (slurp "day2/part2/input.txt")]
    (println (str "The sum of all invalid IDs is: " (solve input-str)))))

(main)

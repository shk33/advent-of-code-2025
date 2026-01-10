
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
  (loop [base 1
         s-base-prev "0"
         invalid-ids []]
    (let [s-base (str base)
          n-invalid (bigint (str s-base s-base))]
      (if (> n-invalid max-id)
        (if (> (count s-base) (count s-base-prev))
          invalid-ids ; Final list
          (recur (inc base) s-base invalid-ids)) ; Continue generating
        (recur (inc base) s-base (conj invalid-ids n-invalid))))))

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
  (let [input-str (slurp "day2/part1/input.txt")]
    (println (str "The sum of all invalid IDs is: " (solve input-str)))))

(main)

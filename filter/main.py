def main() -> None:
    subjects = ["Алгебра", "Биология", "Программирование", "Литература"]
    surnames = ["Смирнов", "Волков", "Морозов", "Лебедев"]
    names = ["Никита", "Елена", "Артем", "София"]
    dates = ["11.04.2026", "12.04.2026", "13.04.2026", "14.04.2026"]

    filter_subject = "Алгебра"
    filter_date = "11.04.2026"

    for subject in subjects:
        for surname in surnames:
            for name in names:
                for date in dates:
                    if subject == filter_subject and date == filter_date:
                        print(subject, surname, name, date)


if __name__ == "__main__":
    main()

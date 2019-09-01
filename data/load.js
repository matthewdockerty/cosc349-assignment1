db = db.getSiblingDB("recipesdb");
db.createCollection("recipes");

use recipesdb;

db.recipes.insert({
	name: "Chicken Milanese with Spaghetti",
	method: "1. Peel and finely slice the garlic. Pick the basil leaves, then finely chop the stalks.<br><br>2. Heat a splash of oil in a medium saucepan over a medium heat. Add the garlic and basil stalks and cook for 2 minutes, or until golden.<br><br>3. Tip the tinned tomatoes into the pan and squash them down with the back of a spoon. Fill the empty tomato tin with water and pour this in too. Season and simmer for 30 minutes, or until reduced, glossy and thick.<br><br>4. Place the chicken breasts on a board and cover with a double layer of clingfilm. Bash with a saucepan to flatten them to 5mm thick.<br><br>5. Tip the flour into one bowl, then crack and beat the eggs in a second bowl. Add the breadcrumbs to a third, finely grate in half of the Parmesan, then shake to combine.<br><br>6. Coat the chicken in the flour, then the egg and finally in the cheesy breadcrumbs, until thoroughly coated.<br><br>7. Heat a lug of oil in a large frying pan over a medium heat and cook the chicken for 3 to 4 mins on each side, or until golden and the meat is cooked through.<br><br>8. Cook the spaghetti according to the packet instructions, then drain and add to the tomato sauce, loosening with a little of the cooking water if needed.<br><br>9. Serve the pasta alongside the crispy chicken, with the remaining Parmesan and basil leaves scattered over.<br><br>(Recipe from https://www.jamieoliver.com/recipes/chicken-recipes/chicken-milanese-with-spaghetti/)",
	ingredients: [
		"2 cloves of garlic",
		"½ a bunch of fresh basil",
		"olive oil",
		"1 x 400 g tin of plum tomatoes",
		"2 x 150 g skinless free-range chicken breasts",
		"100 g plain flour",
		"2 large free-range eggs",
		"100 g breadcrumbs",
		"30 g Parmesan cheese",
		"150 g dried spaghetti"
	],
	image: cat("/vagrant/data/recipe1.b64")
});

db.recipes.insert({
    name: "Apple & Pecan Porridge",
    method: "1. For the basic porridge, place the oats and the milk (or 600ml water) into a large pan over a medium heat, and add a tiny pinch of sea salt.<br><br>2. Bring to a steady simmer for 5 to 6 minutes, stirring often to give you a smooth, creamy porridge, and loosening with extra milk, if needed.<br><br>3. Serve as is, or while it’s blipping away in the pan, follow the next steps to prepare the apple and pecan topping.<br><br>4. Snap the pecans up into little pieces, then toast in a small dry non-stick frying pan over a medium heat for 3 to 4 minutes, or until lightly golden.<br><br> 5. Remove the apple stalk, then use a box grater to coarsely grate it onto a chopping board (core and all).<br><br>6. Stir the grated apple and a little maple syrup through then porridge, then divide between bowls.<br><br>7. Scatter the pecans on top, then drizzle with a little extra maple syrup, if you like.<br><br>(Recipe from https://www.jamieoliver.com/recipes/fruit-recipes/perfect-apple-pecan-porridge/)",
    ingredients: [
        "160 g rolled oats",
        "600 ml milk or organic soya milk",
        "30 g unsalted pecans",
        "1 eating apple",
        "maple syrup"
    ],
    image: cat("/vagrant/data/recipe2.b64")
});

db.recipes.insert({
    name: "Roast Chicken with Potatoes & Carrots",
    method: "1. Preheat the oven to 220°C/425°F/gas 7.<br><br>2. Scrub, trim and halve the carrots lengthways.<br><br>3. Scrub, peel and halve the potatoes, quartering any larger ones. Add to a large roasting tray.<br><br>4. Break the garlic bulb into cloves, leaving them unpeeled, then lightly crush with the flat side of a knife. Pick the rosemary leaves, discarding the stalks. Add the garlic and rosemary leaves to the tray.<br><br>5. Drizzle with oil, season with sea salt and black pepper, then toss well and spread out in an even layer.<br><br>6. Rub the chicken all over with a pinch of salt and pepper and a drizzle of oil. Stuff the chicken cavity with the whole lemon and the thyme sprigs.Place the chicken in the tray, on top of the vegetables.<br><br>7. Reduce the oven temperature to 200ºC/400ºF/gas 6, then add the chicken and roast for 45 minutes.<br><br>8. Carefully remove the tray from the oven, use tongs to turn the vegetables over, then spoon any juices from the tray over the chicken.<br><br>9. Return the tray to the oven for a further 30 minutes, or until the chicken is cooked through. To check, pierce a chicken thigh with the tip of a sharp knife – if the juices run clear, it’s done. Otherwise return the tray to the oven, cook for a little while longer and repeat the test.<br><br>10. Once cooked, transfer the chicken to a board and return the vegetables to the oven for a final 5 minutes to crisp up, if needed.<br><br>11. Cover the chicken with a layer of tin foil and a tea towel, then leave to rest for 10 to 15 minutes.<br><br>12. Using a sharp carving knife, carve up the chicken, then serve with the roasted veg. Delicious with a green salad on the side.<br><br>(Recipe from https://www.jamieoliver.com/recipes/chicken-recipes/roast-chicken-with-potatoes-carrots/)",
    ingredients: [
        "500 g carrots",
        "600 g potatoes",
        "1 bulb of garlic",
        "5 sprigs of fresh rosemary",
        "olive oil",
        "1 x 1.6 kg whole free-range chicken",
        "1 lemon",
        "5 sprigs of fresh thyme"
    ],
    image: cat("/vagrant/data/recipe3.b64")
});

db.recipes.insert({
    name: "Minestrone Soup",
    method: "1. Peel and finely chop the garlic and onion. Trim and roughly chop the carrots, celery and courgette, then add the vegetables to a large bowl.<br><br>2. Cut the ends off the leek, quarter it lengthways, wash it under running water, then cut into 1cm slices. Add to the bowl.<br><br>3. Scrub and dice the potato. Drain the cannellini beans, then set aside. Finely slice the bacon.<br><br>4. Heat 2 tablespoons of oil in a large saucepan over a medium heat. Add the bacon and fry gently for 2 minutes, or until golden.<br><br>5. Add the garlic, onion, carrots, celery, courgette, leek, oregano and bay and cook slowly for about 15 minutes, or until the vegetables have softened, stirring occasionally.<br><br>6. Add the potato, cannellini beans and plum tomatoes, then pour in the vegetable stock. Stir well, breaking up the tomatoes with the back of a spoon.<br><br>7. Cover with a lid and bring everything slowly to the boil, then simmer for about 30 minutes, or until the potato is cooked through. Meanwhile...<br><br>8. Remove and discard any tough stalks bits from the greens, then roughly chop.<br><br>9. Using a rolling pin, bash the pasta into pieces while it’s still in the packet or wrap in a clean tea towel.<br><br>10. To check the potato is cooked, pierce a chunk of it with a sharp knife – if it pierces easily, it’s done.<br><br>11. Add the greens and pasta to the pan, and cook for a further 10 minutes, or until the pasta is al dente. This translates as ‘to the tooth’ and means that it should be soft enough to eat, but still have a bit of a bite and firmness to it. Try some just before the time is up to make sure you cook it perfectly.<br><br>12. Add a splash more stock or water to loosen, if needed.<br><br>13. Pick over the basil leaves (if using) and stir through. Season to taste with sea salt and black pepper, then serve with a grating of Parmesan and a slice of wholemeal bread, if you like.<br><br>(Recipe from https://www.jamieoliver.com/recipes/vegetables-recipes/minestrone-soup/)",
    ingredients: [
        "1 clove of garlic",
        "1 red onion",
        "2 carrots",
        "2 sticks of celery",
        "1 courgette",
        "1 small leek",
        "1 large potato",
        "1 x 400 g tin of cannellini beans",
        "2 rashers of higher-welfare smoked streaky bacon",
        "olive oil",
        "½ teaspoon dried oregano",
        "1 fresh bay leaf",
        "2 x 400 g tins of plum tomatoes",
        "1 litre organic vegetable stock",
        "1 large handful of seasonal greens, such as savoy cabbage, curly kale, chard",
        "100 g wholemeal pasta",
        "½ a bunch of fresh basil, optional",
        "Parmesan cheese"
    ],
    image: cat("/vagrant/data/recipe4.b64")
});

db.recipes.insert({
    name: "Reuben-ish Sandwich",
    method: "1. Grill the slices of bread on a griddle pan until lightly toasted on both sides then spread one side of each with mayonnaise. Put some of the sauerkraut and some of the chilli on 2 of the slices, and top with a couple of slices of pastrami. Top with the remaining sauerkraut and chilli and the sliced gherkins, then grate the Swiss cheese over the top.<br><br>2. Preheat a hot grill. Place the slices with toppings under the grill until the cheese is melted and dribbling.<br><br>3. Stack the sandwich together, adding a few watercress leaves and finishing with the final slice of toast. Press down lightly and use wooden skewers to hold together. Tuck in!<br><br>(Recipe from https://www.jamieoliver.com/recipes/beef-recipes/reuben-ish-sandwich/)",
    ingredients: [
        "2 big slices rye bread, 1cm in size",
        "low fat mayonnaise",
        "3 heaped tablespoons sauerkraut",
        "1 fresh red chilli, deseeded and finely sliced",
        "3 slices pastrami",
        "a few gherkins",
        "60 g Swiss cheese",
        "1 handful watercress leaves, to serve"
    ],
    image: cat("/vagrant/data/recipe5.b64")
});

db.recipes.insert({
    name: "Spiced Sea Bass with Caramelised Fennel",
    method: "1. Reserving the tops, cut the fennel into wedges.<br><br>2. Place the butter and a splash of oil in a pan over a low heat, add the fennel and cook for 40 minutes, or until golden and caramelised.<br><br>3. Preheat the oven to 180ºC/350ºF/gas 4.<br><br>4. Score the fish on each side. Peel and finely chop the garlic, deseed and finely chop the chilli, then rub all over the fish with the lemon zest.<br><br>5. Season with sea salt and black pepper, then slice the lemon into rounds and place 2 to 3 slices inside each fish cavity.<br><br>6. Bake for 20 minutes, or until cooked.<br><br>7. Serve with the fennel, scattered with the reserved tops.<br><br>(Recipe from https://www.jamieoliver.com/recipes/fish-recipes/spiced-sea-bass-with-caramelised-fennel/)",
    ingredients: [
        "3 bulbs of fennel",
        "40 g butter",
        "olive oil",
        "2 whole seabass, gutted and scaled, from sustainable sources",
        "1 clove of garlic",
        "1 fresh red chilli",
        "1 lemon"
    ],
    image: cat("/vagrant/data/recipe6.b64")
});
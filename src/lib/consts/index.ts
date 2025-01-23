const adjectives = [
    // Personality traits
    'Adorable', 'Bubbly', 'Cheerful', 'Dreamy', 'Fluffy', 'Gentle', 'Happy', 'Jolly',
    'Kind', 'Lovely', 'Merry', 'Noble', 'Peppy', 'Quirky', 'Sweet', 'Tender',
    
    // Colors & Nature
    'Golden', 'Rosy', 'Sunny', 'Misty', 'Coral', 'Azure', 'Maple', 'Crystal',
    
    // Food-inspired
    'Honey', 'Sugar', 'Cookie', 'Candy', 'Mochi', 'Berry', 'Vanilla', 'Caramel',
    
    // Magical
    'Sparkly', 'Magical', 'Mystic', 'Starry', 'Wonder', 'Fairy', 'Cosmic', 'Twinkle'
];

const nouns = [
    // Cute animals
    'Kitten', 'Puppy', 'Bunny', 'Panda', 'Penguin', 'Hamster', 'Fox', 'Duckling',
    
    // Food & Sweets
    'Cupcake', 'Cookie', 'Pudding', 'Biscuit', 'Pancake', 'Waffle', 'Sundae', 'Muffin',
    'Butterscotch', 'Donut', 'Marshmallow', 'Chocolate',
    
    // Nature
    'Blossom', 'Cloud', 'Petal', 'Rainbow', 'Sunshine', 'Dewdrop', 'Feather', 'Bubble',
    
    // Magical creatures
    'Unicorn', 'Dragon', 'Phoenix', 'Sprite', 'Pixie', 'Fairy', 'Angel', 'Spirit',
    
    // Celestial
    'Star', 'Moon', 'Comet', 'Aurora', 'Galaxy', 'Nova', 'Nebula', 'Stardust'
];

const avatarIds = ['kazuha', 'diluc', 'ganyu', 'hutao', 'shotgun', 'shenhe'];

export 	const avatars = [
    { id: 'kazuha', src: '/kazuha.png', color: '#ff5144' },
    { id: 'diluc', src: '/diluc.png', color: '#6d2626' },
    { id: 'ganyu', src: '/ganyu.png', color: '#3F51B5' },
    { id: 'hutao', src: '/hutao.png', color: '#a51308' },
    { id: 'shotgun', src: '/shotgun.png', color: '#673AB7' },
    { id: 'shenhe', src: '/shenhe.png', color: '#175d7d' }
];

export function generateRandomUsername(): string {
    const adjective = adjectives[Math.floor(Math.random() * adjectives.length)];
    const noun = nouns[Math.floor(Math.random() * nouns.length)];
    return `${adjective}${noun}`;
}

export function getRandomAvatarId(): string {
    return avatarIds[Math.floor(Math.random() * avatarIds.length)];
} 